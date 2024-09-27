package main

/*
github.com/spf13/cobra - это библиотека для создания CLI приложений на Go.
Она предоставляет простой и удобный интерфейс для создания команд и флагов.
*/
import (
	"github.com/spf13/cobra"
)

/*
Для создания CLI приложения с помощью cobra, необходимо создать корневую команду,
которая будет содержать все остальные команды и флаги.

Параметр Use у рутовой команды используется только для отображения
в справке, и не влияет на логику программы.
Параметр Short используется для краткого описания команды.
Параметр Long используется для подробного описания команды.
Параметр Args используется для валидации аргументов команды.
Параметр Run используется для указания логики выполнения команды.
Более подробно о параметрах команд можно прочитать в документации:
https://pkg.go.dev/github.com/spf13/cobra#Command

В данном примере создается корневая команда rootCmd, которая содержит
команды versionCmd, helloCmd и repeatCmd, а также собственную логику выполнения.

Любая команда может содержать одновременно и дочерние команды, и флаги и собственную логику.

Разрешение команды происходит по следующему приоритету:
1. Флаги
2. Есть ли у команды дочерние команды
3. Есть ли у команды логика выполнения

То есть, если у команды есть флаги, то они будут обработаны в первую очередь.
Если у команды есть дочерняя команда, соответствующая 1-му аргументу, то она будет выполнена.
Если у команды есть логика выполнения, то она будет выполнена только если не было дочерней команды.

Например:
  - go run ./cmd/02-cli version выведет Go on Go v1.0.0
  - go run ./cmd/02-cli hello World выведет
    Go on Go
    Hello
    World
*/
var rootCmd = &cobra.Command{
	Use:   "go-on-go",
	Short: "Go on Go is a Go lang workshop",
	Long:  `Go on Go is a Go lang workshop`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		println("Go on Go")

		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			println("Verbose output")
		}

		for _, arg := range args {
			println(arg)
		}
	},
}

/*
У вложенных команд параметр Use используется для указания имени команды.
*/
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Go on Go",
	Long:  `Print the version number of Go on Go`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetString("verbose")

		if verbose == "true" {
			println("Verbose output")
		}

		println("Go on Go v1.0.0")
	},
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print Hello, World!",
	Long:  `Print Hello, World!`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello,", args[0])
	},
}

/*
RunE - это аналог Run, но с возможностью возвращения ошибки из функции.
Ошибка будет возвращен как результат выполнения Execute().
*/
var repeatCmd = &cobra.Command{
	Use:   "repeat",
	Short: "Print a string multiple times",
	Long:  `Print a string multiple times`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repeat, err := cmd.Flags().GetInt("repeat")

		if err != nil {
			return err
		}

		for i := 0; i < repeat; i++ {
			println(args[0])
		}

		return nil
	},
}

/*
Для создания флагов используется метод Flags() у инстанса команды.
Для создания флага используется методы String, StringP, Int, IntP, Bool, BoolP и т.д.
Подробнее о флагах можно прочитать в документации:
https://pkg.go.dev/github.com/spf13/cobra#Command.Flags

Методы с суффиксом P позволяют указать короткое и длинное имя флага.
*/
func makeRepeatCmd() *cobra.Command {
	repeatCmd.Flags().IntP("repeat", "r", 1, "Number of times to repeat the string")

	return repeatCmd
}

/*
Для добавления команды в родительскую команду используется метод AddCommand.

Execute() запускает выполнение команды и передает управление в соответствующую логику выполнения.
*/
func main() {
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	versionCmd.Flags().StringP("verbose", "v", "false", "Verbose output")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(makeRepeatCmd())

	err := rootCmd.Execute()

	if err != nil {
		println(err.Error())
	}
}
