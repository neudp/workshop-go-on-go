package vanila

import (
	loggingApp "goOnGo/internal/swapi/application/logging"
	"goOnGo/internal/swapi/infrastructure/environment"
	loggingInfra "goOnGo/internal/swapi/infrastructure/logging"
	"goOnGo/internal/swapi/infrastructure/swapi"
	"goOnGo/internal/swapi/infrastructure/transport"
	"goOnGo/internal/swapi/use-case"
	"strconv"
)

/*
По сути dependency injection - это просто сборка объектов из его зависимостей.
*/

type App struct {
	getCharacterHandler *useCase.GetCharacterHandler
}

func NewApp() (*App, error) {
	env, err := environment.Read()
	if err != nil {
		return nil, err
	}

	cfg, err := env.ToConfig()
	if err != nil {
		return nil, err
	}

	logFilter := loggingInfra.NewFilter(cfg.MinLoglevel())
	logWriter := loggingInfra.NewWriter()
	logger := loggingApp.NewLogger(logFilter, logWriter)

	httpClient := transport.NewHttpClient(cfg.SwapiURL(), logger)
	swapiClient := swapi.NewClient(httpClient)
	planetsClient := swapi.NewPlanetsClient(swapiClient, logger)
	charactersClient := swapi.NewCharactersClient(swapiClient, planetsClient, logger)

	return &App{getCharacterHandler: useCase.NewGetCharacterHandler(charactersClient, logger)}, nil
}

func (app *App) Hadle(id string) (*useCase.CharacterDto, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return app.getCharacterHandler.Handle(&useCase.GetCharacterQuery{IdValue: idInt})
}
