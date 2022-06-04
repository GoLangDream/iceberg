package web_test

import (
	"encoding/json"
	"github.com/GoLangDream/iceberg"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/web"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func initWeb() {
	environment.Set(environment.Test)

	web.RouterDraw = func(router *web.Router) {
		router.GET("/get_params/:id", "test_example#get_params")
		router.GET("/get_query", "test_example#get_query")

		router.GET("/has_filter", "test_example#has_filter_action")
		router.GET("/no_filter", "test_example#no_filter_action")

		router.GET("/hello", "home#index")

		router.GET("/set_session", "home#set_session")
		router.GET("/get_session", "home#get_session")

		router.GET("/set_cookie", "home#set_cookie")
		router.GET("/get_cookie", "home#get_cookie")
	}
	web.RegisterController(TestExampleController{})
	iceberg.InitApplication()
}

func TestWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iceberg web模块测试\n")

}

var _ = BeforeSuite(func() {
	initWeb()
})

func getBody(response *http.Response) string {
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func parseBody(response *http.Response, v any) {
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, v)
}
