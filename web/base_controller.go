package web

import (
	"fmt"
	"github.com/gin-contrib/sessions"
)

type BaseController struct {
	name    string
	context *HttpContext

	session sessions.Session
}

func newBaseController(name string, ctx *HttpContext) *BaseController {
	baseController := BaseController{
		name,
		ctx,
		ctx.session(),
	}
	return &baseController
}

func (c *BaseController) Text(body string) {
	if c.context != nil {
		c.context.text(body)
	} else {
		fmt.Println("http context 没有初始化")
	}
}

func (c *BaseController) Session(name string, val ...interface{}) interface{} {
	switch len(val) {
	case 0:
		return c.session.Get(name)
	case 1:
		c.session.Set(name, val[0])
		c.session.Save()
	}
	return nil
}
