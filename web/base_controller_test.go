package web_test

import (
	"github.com/GoLangDream/iceberg/web"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type TestBaseController struct {
	*web.BaseController
}

func (c *TestBaseController) GetParams() {
	id := c.Param("id")
	c.Text("id is " + id)
}

func (c *TestBaseController) GetQuery() {
	name := c.Query("name")
	age := c.Query("age", "1")
	c.Text("name is " + name + " age is " + age)
}

var _ = Describe("BaseController", Ordered, func() {
	var server *web.Server
	var routerDraw = func(router *web.Router) {
		router.GET("/get_params/:id", "test_base#get_params")
		router.GET("/get_query", "test_base#get_query")
	}
	BeforeAll(func() {
		server = web.CreateServer("", routerDraw)
		web.RegisterController(TestBaseController{})

		server.InitServer()
	})

	Context("http提交参数获取", func() {
		It("GET /get_params/123， 应该能得到值为 123 的 id", func() {
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
