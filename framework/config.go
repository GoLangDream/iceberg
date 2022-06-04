package framework

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

func InitConfig(application ApplicationConfig) {
	config.AddDriver(yaml.Driver)
	err := config.LoadExists(configFiles...)
	if err != nil {
		log.Info("加载配置文件错误")
		panic(err)
	}

	log.Infof("加载配置文件 %s\n", strings.Join(configFiles, ", "))

	//for _, configFile := range configFiles {
	//	log.Infof("加载配置文件 [%s]\n", configFile)
	//	configFile := filepath.Join(application.HomePath(), configFile)
	//	config.LoadFiles(configFile)
	//}
}
