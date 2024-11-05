package uberFx

import (
	"go.uber.org/fx"
	getCharacter "goOnGo/internal/swapi/application/get-character"
	loggingApp "goOnGo/internal/swapi/application/logging"
	"goOnGo/internal/swapi/infrastructure/environment"
	loggingInfra "goOnGo/internal/swapi/infrastructure/logging"
	"goOnGo/internal/swapi/infrastructure/swapi"
	"goOnGo/internal/swapi/infrastructure/transport"
	"goOnGo/internal/swapi/model/config"
	"goOnGo/internal/swapi/model/logging"
	"goOnGo/internal/swapi/use-case"
	"strconv"
)

/*
Uber Fx - собирает зависимости непосредственно в рантайме, используя рефлексию.
Рантайм сборка дает некоторую гибкость, но снижает производительность.
Особенностью Uber Fx по сравнению с Google Wire является то, что Uber Fx не отдает
готовый контейнер, а ожидает функцию, которая будет использовать этот контейнер.
*/

func ProvideHttpClient(cfg *config.Config, logger logging.Logger) swapi.Doer {
	return transport.NewHttpClient(cfg.SwapiURL(), logger)
}

func ProvideConfig(env *environment.Environment) (*config.Config, error) {
	return env.ToConfig()
}

func ProvideLoggingFilter(cfg *config.Config) loggingApp.Filter {
	return loggingInfra.NewFilter(cfg.MinLoglevel())
}

type App struct {
	Handler *useCase.GetCharacterHandler
}

func newApp(handler *useCase.GetCharacterHandler) *App {
	return &App{Handler: handler}
}

func (app *App) Handle(id string) (*useCase.CharacterDto, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return app.Handler.Handle(&useCase.GetCharacterQuery{IdValue: idInt})
}

func Do(do func(app *App) error) error {
	app := fx.New(
		fx.Provide(newApp),
		fx.Provide(useCase.NewGetCharacterHandler),
		fx.Provide(fx.Annotate(
			swapi.NewCharactersClient,
			fx.As(new(getCharacter.Repository)),
		)),
		fx.Provide(swapi.NewClient),
		fx.Provide(ProvideHttpClient),
		fx.Provide(ProvideConfig),
		fx.Provide(environment.Read),
		fx.Provide(swapi.NewPlanetsClient),
		fx.Provide(fx.Annotate(
			loggingApp.NewLogger,
			fx.As(new(logging.Logger)),
		)),
		fx.Provide(ProvideLoggingFilter),
		fx.Provide(fx.Annotate(
			loggingInfra.NewWriter,
			fx.As(new(loggingApp.Writer)),
		)),
		fx.Invoke(do),
		fx.NopLogger,
	)

	return app.Err()
}
