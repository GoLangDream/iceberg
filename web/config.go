package web

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func initConfig() {
	config.AddDriver(yaml.Driver)
}
