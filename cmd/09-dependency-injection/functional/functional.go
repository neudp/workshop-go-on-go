package functional

import (
	getCharacter "goOnGo/internal/swapi-func/application/get-character"
	loggingApp "goOnGo/internal/swapi-func/application/logging"
	"goOnGo/internal/swapi-func/infrastructure/environment"
	filterLog "goOnGo/internal/swapi-func/infrastructure/logging"
	"goOnGo/internal/swapi-func/infrastructure/swapi"
	"goOnGo/internal/swapi-func/infrastructure/transport"
	"goOnGo/internal/swapi-func/model/logging"
	"goOnGo/internal/swapi-func/use-case"
	"strconv"
)

func GetCharacter(id string) (*use_case.CharacterDto, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	env, err := environment.Read()
	if err != nil {
		return nil, err
	}

	cfg, err := env.ToConfig()
	if err != nil {
		return nil, err
	}

	logLevel := loggingApp.NewLogLevel(
		filterLog.NewFilterLog(cfg.MinLogLevel()),
		filterLog.NewWriteLog(),
	)
	logger := logging.NewLogger(logLevel)
	doRequest := transport.NewDoRequest(cfg.SwapiURL(), logLevel)
	doGetRequest := swapi.NewDoGetRequest(swapi.DoRequest(doRequest), logger)
	charactersClient := swapi.NewGetCharacter(doGetRequest, logger)

	return use_case.NewGetCharacter(
		getCharacter.Find(charactersClient),
		logger.Info,
	)(&use_case.GetCharacterQuery{IdValue: idInt})
}
