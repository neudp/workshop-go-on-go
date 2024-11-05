package filterLog

import (
	loggingApp "goOnGo/internal/swapi-func/application/logging"
	"goOnGo/internal/swapi-func/model/logging"
)

func NewFilterLog(minLevel logging.Level) loggingApp.FilterLog {
	return func(record *logging.Record) bool {
		return record.Level() >= minLevel
	}
}
