package iceberg

import (
	"github.com/GoLangDream/iceberg/initializers"
	"github.com/GoLangDream/iceberg/log"
	"os"
)

func InitApplication() {
	path, err := os.Getwd()

	if err != nil {
		log.Debugf("运行项目 [%s] 失败", path)
		panic(err)
	} else {
		log.Debugf("运行项目 [%s]", path)
	}

	initializers.Init()
}

func StartApplication() {
	initializers.Start()
}
