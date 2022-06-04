package log

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"time"
)

var green = color.New(color.FgGreen).SprintfFunc()

func init() {
	logrus.SetFormatter(&loggerFormat{})
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

func Output() {

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
