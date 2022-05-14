package web

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type HomeController struct {
	*BaseController
}

func (h *HomeController) GetParams() {
	id := h.Param("id")
	h.Text("id is " + id)
}

func (h *HomeController) GetQuery() {
	name := h.Query("name")
	age := h.Query("age", "1")
	h.Text("name is " + name + " age is " + age)
}

var _ = Describe("BaseController", Ordered, func() {
	BeforeAll(func() {
		RegisterController(HomeController{})

		ApplicationRouterDraw = func(router Router) {
			router.GET("/get_params/:id", "home#get_params")
			router.GET("/get_query", "home#get_query")
		}

		InitServer()
	})

	Context("http提交参数获取", func() {
		It("GET /get_params/123， 应该能得到值为123的id", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/get_params/123", nil)
			server.engine.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("id is 123"))
		})

	})
})
