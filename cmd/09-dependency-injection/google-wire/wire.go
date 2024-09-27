//go:build wireinject
// +build wireinject

package googleWire

import (
	"fmt"
	"github.com/google/wire"
	"goOnGo/internal/swapi/config"
	model2 "goOnGo/internal/swapi/model"
	"goOnGo/internal/swapi/swapi"
	"goOnGo/internal/swapi/transport"
	"goOnGo/internal/swapi/use-case"
	"strconv"
	"strings"
	"time"
)

/*
Wire генерирует провайдеры для ваших структур и функций, чтобы упростить
управление зависимостями в процессе разработки. Wire анализирует код и компилирует
его в провайдеры, которые можно использовать для создания объектов и внедрения
зависимостей.

Однако Wire имеет некоторые ограничения:
- Имена зависимостей должны быть уникальными, нельзя использовать 2 разные
  реализации одного интерфейса
- Wire не поддерживает циклические зависимости (что в целом хорошо)
- Wire несколько усложняет фабричные методы требуя их вычленения в отдельные типы
*/

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

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

type SwapiTransportProvider func() swapi.Transport                   // фабричный метод
type SwapiClientProvider func() *swapi.Swapi                         // фабричный метод
type GetCharacterHandlerProvider func() *useCase.GetCharacterHandler // фабричный метод

type App struct {
	characterHandler GetCharacterHandlerProvider
}

func newApp(characterHandler GetCharacterHandlerProvider) *App {
	return &App{
		characterHandler: characterHandler,
	}
}

func (app *App) GetCharacter(id string) (*model2.Character, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return app.characterHandler().Handle(useCase.NewGetCharacterQuery(idInt))
}

func ProvideSwapiTransport(cfg *config.Config, logger model2.Logger) SwapiTransportProvider {
	return func() swapi.Transport {
		return transport.NewSwapiClient(cfg, logger)
	}
}

func ProvideSwapiClient(transport SwapiTransportProvider, logger model2.Logger) SwapiClientProvider {
	return func() *swapi.Swapi {
		return swapi.New(transport(), logger)
	}
}

func ProvideCharacterHandler(client SwapiClientProvider, logger *Logger) GetCharacterHandlerProvider {
	return func() *useCase.GetCharacterHandler {
		return useCase.NewGetCharacterHandler(client(), logger)
	}
}

func NewApp() (*App, error) {
	panic(wire.Build(
		newApp,
		config.Build,
		NewLogger,
		wire.Bind(new(model2.Logger), new(*Logger)),
		ProvideSwapiTransport,
		ProvideSwapiClient,
		ProvideCharacterHandler,
	))
}
