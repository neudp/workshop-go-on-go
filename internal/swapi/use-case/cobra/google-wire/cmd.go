package googleWire

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"goOnGo/internal/swapi/use-case/cobra/google-wire/wire"
	"goOnGo/internal/swapi/use-case/dto"
	"strconv"
)

type GetCharacterQuery struct {
	IdValue int
}

func (query *GetCharacterQuery) Id() int {
	return query.IdValue
}

var cmd = &cobra.Command{
	Use:   "wire",
	Short: "Swapi application with Google Wire",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		getCharacter, err := wire.NewGetCharacterHandler()
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
