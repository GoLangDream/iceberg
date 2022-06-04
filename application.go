package iceberg

import (
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/framework"
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/web"
	"os"
)

var server *web.Server

func InitApplication(application framework.ApplicationConfig) {
	err := os.Chdir(application.HomePath())
	if err != nil {
		log.Infof("运行项目 [%s] 失败", application.HomePath())
		panic(err)
	} else {
		log.Infof("运行项目 [%s]", application.HomePath())
	}

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
