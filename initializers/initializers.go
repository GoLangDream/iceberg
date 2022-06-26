package initializers

import (
	"github.com/GoLangDream/iceberg/config"
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/web"
	"github.com/GoLangDream/iceberg/work"
)

func Init() {
	environment.Init()
	config.Init()
	log.Init()
	database.Init()
	web.Init()
}

func Start() {
	work.Start()
	web.Start()
}
