package vanilla

import (
	"encoding/json"
	"github.com/spf13/cobra"
	getCharacter "goOnGo/internal/swapi-func/application/get-character"
	loggingApp "goOnGo/internal/swapi-func/application/logging"
	"goOnGo/internal/swapi-func/infrastructure/environment"
	filterLog "goOnGo/internal/swapi-func/infrastructure/logging"
	"goOnGo/internal/swapi-func/infrastructure/swapi"
	"goOnGo/internal/swapi-func/infrastructure/transport"
	"goOnGo/internal/swapi-func/model/logging"
	"goOnGo/internal/swapi-func/use-case/dto"
	"strconv"
)

type GetCharacterQuery struct {
	IdValue int
}

func (query *GetCharacterQuery) Id() int {
	return query.IdValue
}

var cmd = &cobra.Command{
	Use:   "vanilla-func",
	Short: "Swapi application vanilla functional",
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

		logLevel := loggingApp.NewLogLevel(
			filterLog.NewFilterLog(cfg.MinLogLevel()),
			filterLog.NewWriteLog(),
		)
		logger := logging.NewLogger(logLevel)
		doRequest := transport.NewDoRequest(cfg.SwapiURL(), logLevel)
		doGetRequest := swapi.NewDoGetRequest(swapi.DoRequest(doRequest), logger)
		charactersClient := swapi.NewGetCharacter(doGetRequest, logger)

		getCharacter := getCharacter.New(getCharacter.Find(charactersClient), logger.Info)
		if err != nil {
			return err
		}

		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}

		chrctr, err := getCharacter(&GetCharacterQuery{IdValue: id})
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
