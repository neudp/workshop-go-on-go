package vanilla

import (
	"encoding/json"
	"github.com/spf13/cobra"
	getCharacter "goOnGo/internal/swapi/application/get-character"
	loggingApp "goOnGo/internal/swapi/application/logging"
	"goOnGo/internal/swapi/infrastructure/environment"
	loggingInfra "goOnGo/internal/swapi/infrastructure/logging"
	"goOnGo/internal/swapi/infrastructure/swapi"
	"goOnGo/internal/swapi/infrastructure/transport"
	"goOnGo/internal/swapi/use-case/dto"
	"strconv"
)

/*
По сути dependency injection - это просто сборка объектов из его зависимостей.
*/

type GetCharacterQuery struct {
	IdValue int
}

func (query *GetCharacterQuery) Id() int {
	return query.IdValue
}

var cmd = &cobra.Command{
	Use:   "vanilla",
	Short: "Swapi application vanilla DI",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		env, err := environment.Read()
		if err != nil {
			return err
		}

		cfg, err := env.ToConfig()
		if err != nil {
			return err
		}

		logFilter := loggingInfra.NewFilter(cfg.MinLoglevel())
		logWriter := loggingInfra.NewWriter()
		logger := loggingApp.NewLogger(logFilter, logWriter)

		httpClient := transport.NewHttpClient(cfg.SwapiURL(), logger)
		swapiClient := swapi.NewClient(httpClient)
		planetsClient := swapi.NewPlanetsClient(swapiClient, logger)
		charactersClient := swapi.NewCharactersClient(swapiClient, planetsClient, logger)

		getCharacter := getCharacter.NewHandler(charactersClient, logger)
		if err != nil {
			return err
		}

		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}

		chrctr, err := getCharacter.Handle(&GetCharacterQuery{IdValue: id})
		if err != nil {
			panic(err)
		}

		response := &dto.CharacterDto{
			Name:      chrctr.Name(),
			Height:    chrctr.Height(),
			Mass:      chrctr.Mass(),
			HairColor: string(chrctr.HairColor()),
			SkinColor: string(chrctr.SkinColor()),
			EyeColor:  string(chrctr.EyeColor()),
			BirthYear: string(chrctr.BirthYear()),
			Gender:    string(chrctr.Gender()),
			Homeworld: chrctr.Homeworld().Name(),
		}

		resultJson, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return err
		}

		cmd.Println(string(resultJson))

		return nil
	},
}

func Cmd() *cobra.Command {
	return cmd
}
