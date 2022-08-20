package apm

import (
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/log"
	"github.com/gookit/config/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"os"
)

var App *newrelic.Application = nil

func Init() {
	if environment.IsTest() {
		return
	}
	newrelicLicense := config.String("application.newrelic.license", "")
	appName := config.String("application.name") + "_" + environment.Name()
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(newrelicLicense),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Infof("newrelic 加载失败 %s", err)
	} else {
		App = app
		log.Infof("newrelic 加载成功 %s", appName)
	}
}
