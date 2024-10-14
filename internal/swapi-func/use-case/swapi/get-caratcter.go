package swapi

import (
	logingContext "goOnGo/internal/swapi-func/infrastructure/context/loging-context"
	swapiContext "goOnGo/internal/swapi-func/infrastructure/context/swapi-context"
	transportContext "goOnGo/internal/swapi-func/infrastructure/context/transport-context"
	"goOnGo/internal/swapi-func/infrastructure/os"
	"goOnGo/internal/swapi-func/model/domain/character"
	"goOnGo/internal/swapi-func/swapi"
)

/*
Можно заметить, что теперь именно use-case определяет контекст исполнения.
Это гарантирует, что все изменения в контексте будут происходить только по запросу use-case.
А все процессы начавшие исполнение до изменения контекста, будут исполняться в старом контексте.
*/

type GetCharacterQuery struct {
	id int
}

func NewGetCharacterQuery(id int) *GetCharacterQuery {
	return &GetCharacterQuery{id: id}
}

func GetCharacter(query *GetCharacterQuery) (*character.Character, error) {
	cfg, err := os.ConfigFromEnv()

	if err != nil {
		return nil, err
	}

	loggingCtx := logingContext.New(cfg.MinLogLevel())
	transportCtx := transportContext.NewDoSwapiRequestContext(cfg.SwapiURL(), loggingCtx)
	swapiCtx := swapiContext.New(loggingCtx, transportCtx)

	return swapi.GetCharacter(swapiCtx, query.id)
}
