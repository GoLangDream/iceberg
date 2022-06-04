package web_test

import (
	"github.com/GoLangDream/iceberg"
	"github.com/GoLangDream/iceberg/web"
)

type TestExampleController struct {
	*web.BaseController
	beforeActionIsCall int
	afterActionIsCall  int
}

func (c *TestExampleController) Init() {
	c.beforeActionIsCall = 0
	c.afterActionIsCall = 0
	c.BeforeAction(c.setBeforeAction, web.AFH{
		"only": []string{"has_filter_action"},
	})
	c.AfterAction(c.setAfterAction, web.AFH{
		"except": []string{"NoFilterAction"},
	})
}

func (c *TestExampleController) HasFilterAction() {
	c.Json(iceberg.H{
		"before_action_is_call": c.beforeActionIsCall,
		"after_action_is_call":  c.afterActionIsCall,
	})
}

func (c *TestExampleController) NoFilterAction() {
	c.Json(iceberg.H{
		"before_action_is_call": c.beforeActionIsCall,
		"after_action_is_call":  c.afterActionIsCall,
	})
}

func (c *TestExampleController) GetParams() {
	id := c.Param("id")
	c.Text("id is " + id)
}

func (c *TestExampleController) GetQuery() {
	name := c.Query("name")
	age := c.Query("age", "1")
	c.Text("name is " + name + " age is " + age)
}

func (c *TestExampleController) setBeforeAction() {
	c.beforeActionIsCall += 1
}

func (c *TestExampleController) setAfterAction() {
	c.afterActionIsCall += 1
}
