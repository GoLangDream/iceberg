package initializers

import (
	"github.com/GoLangDream/iceberg/config"
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/web"
)

func Init() {
	config.Init()
	database.Init()
	web.Init()
}

func Start() {
	web.Start()
}
