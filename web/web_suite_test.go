package web_test

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	setGinToTestMode()
	RunSpecs(t, "Iceberg web模块测试\n")
}

func setGinToTestMode() {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
}
