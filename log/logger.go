package log

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
)

var green = color.New(color.FgGreen).SprintfFunc()
var prefix = green("[iceberg]")

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&loggerFormat{})
}

func Infof(format string, v ...any) {
	logrus.Infof(format, v...)
}

func Info(v ...any) {
	logrus.Info(v...)
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
	b.WriteString(fmt.Sprintf("%s %s", prefix, entry.Message))
	return b.Bytes(), nil
}
