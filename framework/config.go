package framework

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"path/filepath"
)

var configFiles = []string{"config/database.yml"}

func InitConfig(application ApplicationConfig) {
	config.AddDriver(yaml.Driver)
	for _, configFile := range configFiles {
		fmt.Printf("加载配置文件 [%s]\n", configFile)
		configFile := filepath.Join(application.HomePath(), configFile)
		config.LoadFiles(configFile)
	}
}
