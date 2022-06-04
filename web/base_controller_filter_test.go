package web_test

import (
	"github.com/GoLangDream/iceberg/web"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("BaseController", func() {

	Context("访问没有filter的action", func() {
		It("GET /no_filter， before action 和 after action 都应该不会被调用", func() {
			req := httptest.NewRequest("GET", "/no_filter", nil)
			rep, _ := web.Test(req, 1)
			checkFilterResponse(rep, 0, 0)
		})

		It("GET /has_filter， before action 和 after action 应该会被调用1次", func() {
			req := httptest.NewRequest("GET", "/has_filter", nil)
			rep, _ := web.Test(req)
			checkFilterResponse(rep, 1, 1)
		})

	})

})

func checkFilterResponse(response *http.Response, beforeActionIsCall, afterActionIsCall int) {
	var result map[string]int
	Expect(response.StatusCode).To(Equal(200))
	parseBody(response, &result)
	beforeAction, _ := result["before_action_is_call"]
	afterAction, _ := result["after_action_is_call"]
	Expect(beforeAction).To(Equal(beforeActionIsCall))
	Expect(afterAction).To(Equal(afterActionIsCall))
}
