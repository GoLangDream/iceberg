package config

import (
	"github.com/GoLangDream/iceberg/log"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"strings"
)

var configFiles = []string{
	"config/database.yml",
	"config/application.yml",
}

func Init() {
	config.AddDriver(yaml.Driver)
	err := config.LoadExists(configFiles...)
	if err != nil {
		log.Info("加载配置文件错误")
		panic(err)
	}

	log.Debugf("加载配置文件 %s", strings.Join(configFiles, ", "))
}
