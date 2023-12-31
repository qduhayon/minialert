package logger

import (
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type ServerFormatter struct {
	log.TextFormatter
}

type ClientFormatter struct {
	log.TextFormatter
}

func (f *ServerFormatter) Format(entry *log.Entry) ([]byte, error) {
	var levelColor int

	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 37 // gray
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("\x1b[%dm[%s][%s] %s\x1b[0m\n", levelColor, entry.Time.Format(time.DateTime), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func (f *ClientFormatter) Format(entry *log.Entry) ([]byte, error) {
	var levelColor int

	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 37 // gray
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("\x1b[%dm[%s] %s\x1b[0m\n", levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}
