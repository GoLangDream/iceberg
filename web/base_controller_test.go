package web_test

import (
	"github.com/GoLangDream/iceberg/web"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type BaseTestController struct {
	*web.BaseController
}

func (h *BaseTestController) GetParams() {
	id := h.Param("id")
	h.Text("id is " + id)
}

func (h *BaseTestController) GetQuery() {
	name := h.Query("name")
	age := h.Query("age", "1")
	h.Text("name is " + name + " age is " + age)
}

type BaseTestApplication struct {
}

func (app *BaseTestApplication) RouterDraw(router *web.Router) {
	router.GET("/get_params/:id", "base_test#get_params")
	router.GET("/get_query", "base_test#get_query")
}

var _ = Describe("BaseController", Ordered, func() {
	var server *web.Server

	BeforeAll(func() {
		server = web.CreateServer(&BaseTestApplication{})
		web.RegisterController(BaseTestController{})

		server.InitServer()
	})

	Context("http提交参数获取", func() {
		It("GET /get_params/123， 应该能得到值为123的id", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/get_params/123", nil)
			server.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("id is 123"))
		})

		It("GET /get_query?name=jim&age=23, 能正确读取 name 和 age的值", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/get_query?name=jim&age=23", nil)
			server.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("name is jim age is 23"))
		})
	})

})
