package main

import (
	"github.com/spf13/cobra"
	vanillaFunc "goOnGo/internal/swapi-func/use-case/cobra/vanilla"
	googleWire "goOnGo/internal/swapi/use-case/cobra/google-wire"
	uberFx "goOnGo/internal/swapi/use-case/cobra/uber-fx"
	"goOnGo/internal/swapi/use-case/cobra/vanilla"
)

/*
Dependency injection - это процесс предоставления зависимостей объекту.

В Go это можно сделать с помощью структур, в которых зависимости являются полями.
Или с помощью интерфейсов, которые описывают требуемые методы. Поскольку в Go
не существует наследования и как следствие полиморфизма, то способ
с использованием структур является более понятным, но, как и в классических ООП
языках, это может привести к проблеме с циклическими зависимостями и сложности
в тестировании.

Мы рассмотрим пример 3 способов реализации dependency injection:
- Ванильный метод ручного создания зависимостей
- Google Wire - библиотека от Google, по сути автоматизирующая ванильный метод,
  но с некоторыми ограничениями
- Uber FX - библиотека от Uber, так же основанная на рефлексии, но работающая
  с функциями вместо структур
- Функциональный метод (тот же ванильный метод, но с использованием функционального стиля)
*/

func main() {
	rootCmd := &cobra.Command{
		Use:   "dependency-injection",
		Short: "Dependency injection examples",
	}

	rootCmd.AddCommand(
		googleWire.Cmd(),
		uberFx.Cmd(),
		vanilla.Cmd(),
		vanillaFunc.Cmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
