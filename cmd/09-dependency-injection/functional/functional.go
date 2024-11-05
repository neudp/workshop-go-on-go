package functional

import (
	"goOnGo/internal/swapi-func/application/logging"
	"goOnGo/internal/swapi-func/infrastructure/environment"
	filterLog "goOnGo/internal/swapi-func/infrastructure/filter-log"
	"goOnGo/internal/swapi-func/infrastructure/transport"
	writeLog "goOnGo/internal/swapi-func/infrastructure/write-log"
	"goOnGo/internal/swapi-func/use-case/swapi"
	"strconv"
)

func GetCharacter(id string) (*swapi.CharacterDto, error) {
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

	logLevel := logging.NewLogLevel(
		filterLog.NewFilterLog(cfg.MinLogLevel()),
		writeLog.NewWriteLog(),
	)

	return swapi.NewGetCharacter(
		logLevel,
		transport.NewDoSwapiRequest(logLevel, cfg.SwapiURL()),
	)(&swapi.GetCharacterQuery{Id: idInt})
}
