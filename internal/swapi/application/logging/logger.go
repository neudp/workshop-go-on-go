package logging

import "goOnGo/internal/swapi/model/logging"

type Filter interface {
	Filter(record *logging.Record) bool
}

type Writer interface {
	Write(record *logging.Record)
}

type Logger struct {
	filter Filter
	writer Writer
}

func NewLogger(filter Filter, writer Writer) *Logger {
	return &Logger{
		filter: filter,
		writer: writer,
	}
}

func (logger *Logger) Log(level logging.Level, message string, labels ...*logging.Label) {
	record := logging.NewRecord(level, message, labels...)

	if logger.filter.Filter(record) {
		logger.writer.Write(record)
	}
}

func (logger *Logger) Info(message string, labels ...*logging.Label) {
	logger.Log(logging.Info, message, labels...)
}

func (logger *Logger) Error(message string, labels ...*logging.Label) {
	logger.Log(logging.Error, message, labels...)
}
