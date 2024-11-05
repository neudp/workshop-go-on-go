package logging

import "goOnGo/internal/swapi/model/logging"

type Filter struct {
	minLogLevel logging.Level
}

func NewFilter(minLogLevel logging.Level) *Filter {
	return &Filter{minLogLevel: minLogLevel}
}

func (filter *Filter) Filter(record *logging.Record) bool {
	return record.Level() >= filter.minLogLevel
}
