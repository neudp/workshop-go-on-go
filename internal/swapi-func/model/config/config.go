package config

import "goOnGo/internal/swapi-func/model/logging"

type Config struct {
	swapiURL    string
	minLogLevel logging.Level
}

func New(
	swapiURL string,
	minLogLevel logging.Level,
) *Config {
	return &Config{
		swapiURL:    swapiURL,
		minLogLevel: minLogLevel,
	}
}

func (config *Config) SwapiURL() string {
	return config.swapiURL
}

func (config *Config) MinLogLevel() logging.Level {
	return config.minLogLevel
}

/*
Функциональный подход отдает предпочтение созданию новых объектов вместо изменения существующих.
Это реализует одно из важных правил функционального программирования - чистоту функций.
Чистая функция - это функция, которая не изменяет состояние программы и не зависит от состояния
программы. Таким образом, функция всегда возвращает одинаковый результат при одинаковых входных данных.
*/

func (config *Config) Clone() *Config {
	return &Config{
		swapiURL:    config.swapiURL,
		minLogLevel: config.minLogLevel,
	}
}

/*
Как следствие, у нас нет методов для изменения полей структуры. Мы есть функции, которые создают новый объект
*/

func ChangeMinLogLevel(config *Config, minLogLevel logging.Level) *Config {
	newConfig := config.Clone()
	newConfig.minLogLevel = minLogLevel

	return newConfig
}

/*
Есть лишь одно исключение, когда мы допускаем изменение состояния объекта - исключительно синхронное исполнение
Если структура используется в однопоточном режиме, то можно изменять состояние объекта, но функция обязана возвращать
ссылку

Однако, это метод оптимизации и в функциональной парадигме структуры должны быть неизменяемыми по умолчанию
*/

func ChangeMinLogLevelSync(config *Config, minLogLevel logging.Level) *Config {
	config.minLogLevel = minLogLevel

	return config
}
