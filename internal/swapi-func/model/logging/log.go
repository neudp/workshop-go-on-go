package logging

type LogLevel = func(level Level, message string, labels ...*Label)
type Log = func(message string, labels ...*Label)

func NewLog(logLevel LogLevel, level Level) Log {
	return func(message string, labels ...*Label) {
		logLevel(level, message, labels...)
	}
}

type Logger struct {
	logError Log
	logInfo  Log
}

func NewLogger(logLevel LogLevel) *Logger {
	return &Logger{
		logError: NewLog(logLevel, Error),
		logInfo:  NewLog(logLevel, Info),
	}
}

func (logger Logger) Info(message string, labels ...*Label) {
	logger.logInfo(message, labels...)
}

func (logger Logger) Error(message string, labels ...*Label) {
	logger.logError(message, labels...)
}
