package iceberg

import (
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/framework"
	"github.com/GoLangDream/iceberg/web"
)

var server *web.Server

func InitApplication(application framework.ApplicationConfig) {
	server = web.CreateServer(
		application.HomePath(),
		application.RouterDraw(),
	)
	framework.InitConfig(application)
	database.InitDatabase()
}

func StartApplication() {
	server.Start()
}
