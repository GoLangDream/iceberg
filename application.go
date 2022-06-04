package iceberg

import (
	"github.com/GoLangDream/iceberg/initializers"
	"github.com/GoLangDream/iceberg/log"
	"os"
)

func InitApplication() {
	path, err := os.Getwd()

	if err != nil {
		log.Infof("运行项目 [%s] 失败", path)
		panic(err)
	} else {
		log.Infof("运行项目 [%s]", path)
	}

	initializers.Init()
}

func StartApplication() {
	initializers.Start()
}
