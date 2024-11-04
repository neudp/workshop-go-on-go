package logging

import (
	"goOnGo/internal/swapi-func/model/logging"
)

/*
Используя функциональный подход, мы описываем логику в виде функций, а зависимости передаем через аргументы.
Аргументы-зависимости - это так же функции, выполняющие какую-то единицу работы.
Такой подход называется композицией функций, а сами функции, реализующие композицию, называются функциями высшего порядка.

Как и в объектно-ориентированном подходе, функциональный подход позволяет инкаспулировать логику, но делает это иначе.
Вместо того чтобы описывать интерфейсы, мы описываем функцию без зависимостей как тип, а затем реализуем ее в виде
последовательного вызова или карирования.

//
В данном случае, у нас есть функция Log, для ее работы необходимы функции FilterLog, WriteLog и параметр Level,
то есть полная сигнатура функции Log выглядит так:
func Log(
    // зависимости
	filter func(record *logging.Record) bool
	write func(record *logging.Record)
	// параметры
	level logging.Level,
	// аргументы
	message string,
	labels ...*logging.Label,
)
*/

// Выделим типы

type FilterLog = func(record *logging.Record) bool
type WriteLog = func(record *logging.Record)

// Мы разделелили функцию Log на две функции: Параметризованную и Непараметризованную.

//type LogLevel func(level logging.Level, message string, labels ...*logging.Label)
//type Log func(message string, labels ...*logging.Label)

// Опишем логику непараметризованной функции

func NewLogLevel(filter FilterLog, write WriteLog) logging.LogLevel {
	return func(level logging.Level, message string, labels ...*logging.Label) {
		record := logging.NewRecord(level, message, labels...)

		if filter(record) {
			write(record)
		}
	}
}

// Опишем логику параметризованной функции

//func NewLog(logLevel logging.LogLevel, level logging.Level) logging.Log {
//	return func(message string, labels ...*logging.Label) {
//		logLevel(level, message, labels...)
//	}
//}

/*
Теперь, у нас есть api для создания функций логирования

logInfo := NewLog(NewLogLevel(FilterLogInfo, WriteLog), logging.Info)

logInfo("Hello, World!")

Поскольку функции Log, LogLevel и NewLog не имеют зависимостей, их можно вынести
в отдельный пакет и использовать в других пакетах.
*/
