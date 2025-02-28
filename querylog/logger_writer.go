package querylog

import (
	"reflect"
	"strings"

	"github.com/0xERR0R/blocky/log"
	"github.com/sirupsen/logrus"
)

const loggerPrefixLoggerWriter = "query"

type LoggerWriter struct {
	logger *logrus.Entry
}

func NewLoggerWriter() *LoggerWriter {
	return &LoggerWriter{logger: log.PrefixedLog(loggerPrefixLoggerWriter)}
}

func (d *LoggerWriter) Write(entry *LogEntry) {
	d.logger.WithFields(
		logrus.Fields{
			"cli_ip":  entry.ClientIP,
			"reason":  entry.ResponseReason,
			"rcode":   entry.ResponseCode,
			"qname":   entry.QuestionName,
			"qtype":   entry.QuestionType,
			"answer":  entry.Answer,
			"time_ms": entry.DurationMs,
		},
	).Infof("")
}

func (d *LoggerWriter) CleanUp() {
	// Nothing to do
}

func LogEntryFields(entry *LogEntry) logrus.Fields {
	return withoutZeroes(logrus.Fields{
		"client_ip":       entry.ClientIP,
		"client_names":    strings.Join(entry.ClientNames, "; "),
		"response_reason": entry.ResponseReason,
		"response_type":   entry.ResponseType,
		"response_code":   entry.ResponseCode,
		"question_name":   entry.QuestionName,
		"question_type":   entry.QuestionType,
		"answer":          entry.Answer,
		"duration_ms":     entry.DurationMs,
		"instance":        entry.BlockyInstance,
	})
}

func withoutZeroes(fields logrus.Fields) logrus.Fields {
	for k, v := range fields {
		if reflect.ValueOf(v).IsZero() {
			delete(fields, k)
		}
	}

	return fields
}
