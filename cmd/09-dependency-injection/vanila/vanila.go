package vanila

import (
	"fmt"
	"goOnGo/internal/swapi/config"
	"goOnGo/internal/swapi/model"
	"goOnGo/internal/swapi/swapi"
	"goOnGo/internal/swapi/transport"
	"goOnGo/internal/swapi/use-case"
	"strconv"
	"strings"
	"time"
)

/*
По сути dependency injection - это просто сборка объектов из его зависимостей.
В данном случае мы создаем объект App, который является контейнером для всех
частей приложения. В нашем случае это конфигурация и клиент для обращения к API.

Конфигурация создается в конструкторе NewApp, а клиент создается в методе SwapiClient.
Такой подход обусловлен тем, что конфигурация - это общая для всего приложения и не
изменяется в процессе работы, а клиент может быть создан с разными параметрами и необходим
только во время выполнения запросов к API. Создание конфигурации - достаточно дорогостоящая
операция, и, поскольку она не изменяется, то нет смысла создавать ее каждый раз. Клиент же
может хранить промежуточные данные, контекст, куки и т.д., поэтому его создание каждый раз
проще и безопаснее.
*/

type Logger struct {
	enabled bool
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if !logger.enabled {
		return
	}

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Printf(time.Now().Format("2006-01-02 15-04-05.000")+" INFO: "+format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if !logger.enabled {
		return
	}

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Printf(time.Now().Format("2006-01-02 15-04-05.000")+" ERROR: "+format, args...)
}

type App struct {
	config *config.Config
	logger *Logger
}

func NewApp(log bool) (*App, error) {
	logger := &Logger{enabled: log}
	cfg, err := config.Build(logger)

	if err != nil {
		return nil, err
	}

	return &App{
		config: cfg,
		logger: logger,
	}, nil
}

func (app *App) SwapiClient() *swapi.Swapi {
	client := transport.NewSwapiClient(app.config, app.logger)
	return swapi.New(client, app.logger)
}

func (app *App) GetCharacter(id string) (*model.Character, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	handler := useCase.NewGetCharacterHandler(app.SwapiClient(), app.logger)
	return handler.Handle(useCase.NewGetCharacterQuery(idInt))
}
