package log

import (
	"bytes"
	"fmt"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

var green = color.New(color.FgGreen).SprintfFunc()

type Level = logrus.Level

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

func Init() {
	logrus.SetFormatter(&loggerFormat{})
	if environment.IsTest() {
		logrus.SetOutput(ioutil.Discard)
	}
}

func SetLevel(level Level) {
	logrus.SetLevel(level)
}

func Prefix() string {
	now := time.Now().Format("15:04:05")
	return green(fmt.Sprintf("[iceberg %s]", now))
}

func Infof(format string, v ...any) {
	logrus.Infof(format, v...)
}

func Info(v ...any) {
	logrus.Info(v...)
}

func Debugf(format string, v ...any) {
	logrus.Debugf(format, v...)
}

type loggerFormat struct {
}

func (l *loggerFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(fmt.Sprintf("%s %s", Prefix(), entry.Message))
	b.WriteByte('\n')
	return b.Bytes(), nil
}
