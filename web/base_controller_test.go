package web_test

import (
	"github.com/GoLangDream/iceberg/web"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("BaseController", func() {

	Context("http提交参数获取", func() {
		It("GET /get_params/123， 应该能得到值为 123 的 id", func() {
			req := httptest.NewRequest("GET", "/get_params/123", nil)
			rep, _ := web.Test(req)
			Expect(rep.StatusCode).To(Equal(200))
			Expect(getBody(rep)).To(Equal("id is 123"))
		})

		It("GET /get_query?name=jim&age=23, 能正确读取 name 和 age的值", func() {
			req, _ := http.NewRequest("GET", "/get_query?name=jim&age=23", nil)
			rep, _ := web.Test(req)
			Expect(rep.StatusCode).To(Equal(200))
			Expect(getBody(rep)).To(Equal("name is jim age is 23"))
		})
	})

})
