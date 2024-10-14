package logging

import "time"

type Record struct {
	timestamp time.Time
	level     Level
	message   string
	labels    *Labels
}

func NewRecord(level Level, message string, labels ...*Label) *Record {
	return &Record{
		timestamp: time.Now(),
		level:     level,
		message:   message,
		labels:    NewLabels(labels...),
	}
}

func (r *Record) Timestamp() time.Time {
	return r.timestamp
}

func (r *Record) Level() Level {
	return r.level
}

func (r *Record) Message() string {
	return r.message
}

func (r *Record) Labels() *Labels {
	return r.labels //
}

func (r *Record) Clone() *Record {
	return &Record{
		timestamp: r.timestamp,
		level:     r.level,
		message:   r.message,
		labels:    r.labels.Clone(),
	}
}

func AddLabel(record *Record, label *Label) *Record {
	newRecord := record.Clone()
	newRecord.labels.labels = append(newRecord.labels.labels, label)
	newRecord.labels.index[label.Key()] = append(newRecord.labels.index[label.Key()], label)

	return newRecord
}
