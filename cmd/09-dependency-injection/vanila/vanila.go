package vanila

import (
	"fmt"
	"goOnGo/cmd/09-dependency-injection/config"
	"goOnGo/cmd/09-dependency-injection/model"
	"goOnGo/cmd/09-dependency-injection/swapi"
	"goOnGo/cmd/09-dependency-injection/transport"
	useCase "goOnGo/cmd/09-dependency-injection/use-case"
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

type Logger struct{}

func (l *Logger) Infof(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Printf(time.Now().Format("2006-01-02 15-04-05.000")+" INFO: "+format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Printf(time.Now().Format("2006-01-02 15-04-05.000")+" ERROR: "+format, args...)
}

type App struct {
	config *config.Config
	logger *Logger
}

func NewApp() (*App, error) {
	logger := &Logger{}
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
