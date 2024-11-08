//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	getCharacter "goOnGo/internal/swapi/application/get-character"
	loggingApp "goOnGo/internal/swapi/application/logging"
	"goOnGo/internal/swapi/infrastructure/environment"
	loggingInfra "goOnGo/internal/swapi/infrastructure/logging"
	"goOnGo/internal/swapi/infrastructure/swapi"
	"goOnGo/internal/swapi/infrastructure/transport"
	"goOnGo/internal/swapi/model/config"
	"goOnGo/internal/swapi/model/logging"
)

/*
Wire генерирует провайдеры для ваших структур и функций, чтобы упростить
управление зависимостями в процессе разработки. Wire анализирует код и компилирует
его в провайдеры, которые можно использовать для создания объектов и внедрения
зависимостей.

Однако Wire имеет некоторые ограничения:
- Имена зависимостей должны быть уникальными, нельзя использовать 2 разные
  реализации одного интерфейса
- Wire не поддерживает циклические зависимости (что в целом хорошо)
- Wire несколько усложняет фабричные методы требуя их вычленения в отдельные типы
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

func NewGetCharacterHandler() (*getCharacter.Handler, error) {
	panic(wire.Build(
		getCharacter.NewHandler,   // реализация хендлера
		swapi.NewCharactersClient, // реализация репозитория
		wire.Bind(new(getCharacter.Repository), new(*swapi.CharactersClient)), // связываем интерфейс и реализацию
		swapi.NewClient,        // исполнитель запросов для swapi клиента
		ProvideHttpClient,      // <- провайдер исполнителя HTTP запросов, поскольку конструктор требует строчный аргумент
		ProvideConfig,          // <- провайдер конфига поскольку конфиг строится из окружения
		environment.Read,       // Окружение
		swapi.NewPlanetsClient, // реализация клиента планет
		loggingApp.NewLogger,   // логгер
		wire.Bind(new(logging.Logger), new(*loggingApp.Logger)), // связываем логгер с его реализацией
		ProvideLoggingFilter,   // <- провайдер фильтра для логгера, поскольку конструктор требует скалярный аргумент
		loggingInfra.NewWriter, // реализация врайтера для логгера
		wire.Bind(new(loggingApp.Writer), new(*loggingInfra.Writer)), // связываем врайтер с его реализацией
	))
}
