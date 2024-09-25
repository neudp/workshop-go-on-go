package uberFx

import (
	"fmt"
	"go.uber.org/fx"
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
Uber Fx - собирает зависимости непосредственно в рантайме, используя рефлексию.
Рантайм сборка дает некоторую гибкость, но снижает производительность.
Особенностью Uber Fx по сравнению с Google Wire является то, что Uber Fx не отдает
готовый контейнер, а ожидает функцию, которая будет использовать этот контейнер.
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

func ProvideSwapiTransport(cfg *config.Config, logger model.Logger) SwapiTransportProvider {
	return func() swapi.Transport {
		return transport.NewSwapiClient(cfg, logger)
	}
}

func ProvideSwapiClient(transport SwapiTransportProvider, logger model.Logger) SwapiClientProvider {
	return func() *swapi.Swapi {
		return swapi.New(transport(), logger)
	}
}

func ProvideCharacterHandler(client SwapiClientProvider, logger model.Logger) GetCharacterHandlerProvider {
	return func() *useCase.GetCharacterHandler {
		return useCase.NewGetCharacterHandler(client(), logger)
	}
}

type App struct {
	getCharacterHandler GetCharacterHandlerProvider
}

func NewApp(characterHandler GetCharacterHandlerProvider) *App {
	return &App{
		getCharacterHandler: characterHandler,
	}
}

func (app *App) GetCharacter(id string) (*model.Character, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return app.getCharacterHandler().Handle(useCase.NewGetCharacterQuery(idInt))
}

func Run(do func(*App) error) error {
	app := fx.New(
		fx.Provide(fx.Annotate(
			NewLogger,
			fx.As(new(model.Logger)),
		)),
		fx.Provide(config.Build),
		fx.Provide(ProvideSwapiTransport),
		fx.Provide(ProvideSwapiClient),
		fx.Provide(ProvideCharacterHandler),
		fx.Provide(NewApp),
		fx.Invoke(do),
	)

	return app.Err()
}
