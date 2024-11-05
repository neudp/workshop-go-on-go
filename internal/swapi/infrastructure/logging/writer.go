package logging

import (
	"encoding/json"
	"fmt"
	"goOnGo/internal/swapi/model/logging"
)

type recordDto struct {
	Timestamp string            `json:"timestamp"`
	Level     string            `json:"level"`
	LevelCode int               `json:"level_code"`
	Message   string            `json:"message"`
	Labels    map[string]string `json:"labels"`
}

type Writer struct{}

func NewWriter() *Writer {
	return &Writer{}
}

func (*Writer) levelName(level logging.Level) string {
	switch level {
	case logging.Info:
		return "INFO"
	case logging.Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func (*Writer) levelCode(level logging.Level) int {
	return int(level)
}

func (writer *Writer) Write(record *logging.Record) {
	labels := make(map[string]string)

	for _, key := range record.Labels().All() {
		labels[key] = record.Labels().Get(key)[0].Value()
	}

	dto := recordDto{
		Timestamp: record.Timestamp().Format("2006-01-02 15:04:05.000"),
		Level:     writer.levelName(record.Level()),
		LevelCode: writer.levelCode(record.Level()),
		Message:   record.Message(),
		Labels:    labels,
	}

	bytes, err := json.Marshal(dto)

	if err != nil {
		panic(err) // should never happen
	}

	fmt.Println(string(bytes))
}
