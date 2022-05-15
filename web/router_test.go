package web_test

import (
	"github.com/GoLangDream/iceberg/web"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func checkRouter(method, path, structName, structMethod string, server *web.Server) (result bool) {
	result = false
	for _, info := range server.AllRoutes() {
		if info.Method == method &&
			info.Path == path &&
			structName == info.StructName &&
			structMethod == info.StructMethod {
			result = true
		}
	}
	return
}

type RouterTestApplication struct {
}

func (app *RouterTestApplication) RouterDraw(router *web.Router) {
	router.GET("/hello", "home#index")

	router.GET("/set_session", "home#set_session")
	router.GET("/get_session", "home#get_session")

	router.GET("/set_cookie", "home#set_cookie")
	router.GET("/get_cookie", "home#get_cookie")
}

var _ = Describe("Router", Ordered, func() {
	var server *web.Server
	BeforeAll(func() {
		server = web.CreateServer(&RouterTestApplication{})

		server.InitServer()
	})

	Context("路由", func() {
		It("GET /hello， home#index 会正确的路由到HomeController struct的 Index 方法上", func() {
			Expect(checkRouter(
				"GET",
				"/hello",
				"HomeController",
				"Index",
				server,
			)).To(Equal(true))
		})

		It("GET /set_session， home#index 会正确的路由到HomeController struct的 SetSession 方法上", func() {
			Expect(checkRouter(
				"GET",
				"/set_session",
				"HomeController",
				"SetSession",
				server,
			)).To(Equal(true))
		})
	})
})
