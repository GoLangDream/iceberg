package web_test

import (
	"github.com/GoLangDream/iceberg/web"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type TestFilterController struct {
	*web.BaseController
	beforeActionIsCall int
	afterActionIsCall  int
}

func (c *TestFilterController) Init() {
	c.beforeActionIsCall = 0
	c.afterActionIsCall = 0
	c.BeforeAction(c.setBeforeAction, web.AFH{
		"only": []string{"has_filter_action"},
	})
	c.AfterAction(c.setAfterAction, web.AFH{
		"except": []string{"NoFilterAction"},
	})
}

func (c *TestFilterController) HasFilterAction() {
	c.Json(web.H{
		"before_action_is_call": c.beforeActionIsCall,
		"after_action_is_call":  c.afterActionIsCall,
	})
}

func (c *TestFilterController) NoFilterAction() {
	c.Json(web.H{
		"before_action_is_call": c.beforeActionIsCall,
		"after_action_is_call":  c.afterActionIsCall,
	})
}

func (c *TestFilterController) setBeforeAction() {
	c.beforeActionIsCall += 1
}

func (c *TestFilterController) setAfterAction() {
	c.afterActionIsCall += 1
}

var _ = Describe("BaseController", Ordered, func() {
	var server *web.Server
	var routerDraw = func(router *web.Router) {
		router.GET("/has_filter", "test_filter#has_filter_action")
		router.GET("/no_filter", "test_filter#no_filter_action")
	}
	BeforeAll(func() {
		server = web.CreateServer("", routerDraw)
		web.RegisterController(TestFilterController{})

		server.InitServer()
	})

	Context("访问没有filter的action", func() {
		It("GET /no_filter， before action 和 after action 都应该不会被调用", func() {
			req := httptest.NewRequest("GET", "/no_filter", nil)
			rep, _ := server.Test(req, 1)
			checkFilterResponse(rep, 0, 0)
		})

		It("GET /has_filter， before action 和 after action 应该会被调用1次", func() {
			req := httptest.NewRequest("GET", "/has_filter", nil)
			rep, _ := server.Test(req, 1)
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
