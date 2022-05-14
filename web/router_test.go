package web

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func checkRouter(method, path, structName, structMethod string) (result bool) {
	result = false
	for _, info := range Routes() {
		if info.Method == method &&
			info.Path == path &&
			structName == info.StructName &&
			structMethod == info.StructMethod {
			result = true
		}
	}
	return
}

var _ = Describe("Router", Ordered, func() {
	BeforeAll(func() {
		ApplicationRouterDraw = func(router Router) {
			router.GET("/hello", "home#index")

			router.GET("/set_session", "home#set_session")
			router.GET("/get_session", "home#get_session")

			router.GET("/set_cookie", "home#set_cookie")
			router.GET("/get_cookie", "home#get_cookie")
		}

		InitServer()
	})

	Context("路由", func() {
		It("GET /hello， home#index 会正确的路由到HomeController struct的 Index 方法上", func() {
			Expect(checkRouter(
				"GET",
				"/hello",
				"HomeController",
				"Index",
			)).To(Equal(true))
		})

		It("GET /set_session， home#index 会正确的路由到HomeController struct的 SetSession 方法上", func() {
			Expect(checkRouter(
				"GET",
				"/set_session",
				"HomeController",
				"SetSession",
			)).To(Equal(true))
		})
	})
})
