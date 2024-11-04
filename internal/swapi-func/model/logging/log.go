package logging

type LogLevel = func(level Level, message string, labels ...*Label)
type Log = func(message string, labels ...*Label)

func NewLog(logLevel LogLevel, level Level) Log {
	return func(message string, labels ...*Label) {
		logLevel(level, message, labels...)
	}
}
