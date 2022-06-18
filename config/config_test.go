package config_test

import (
	initConfig "github.com/GoLangDream/iceberg/config"
	"github.com/gookit/config/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Env", func() {

	It("能正确的加载数据库的配置", func() {
		pwd, _ := os.Getwd()
		os.Chdir("../testdata")

		initConfig.Init()

		Expect(config.String("database.adapter")).To(Equal("mysql"))
		Expect(config.String("database.encoding")).To(Equal("utf8mb4"))
		Expect(config.String("database.name")).To(Equal("test_db"))
		Expect(config.String("database.username")).To(Equal("root"))
		Expect(config.String("database.password")).To(Equal("123456"))
		Expect(config.String("database.host")).To(Equal("192.168.2.22"))
		Expect(config.String("database.port")).To(Equal("3306"))

		os.Chdir(pwd)

	})

})
