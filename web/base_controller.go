package web

import (
	"fmt"
	"github.com/GoLangDream/rgo/option"
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

func (c *BaseController) Cookie(name string, val ...interface{}) interface{} {
	valLen := len(val)

	switch {
	case valLen == 0:
		return c.context.cookie(name)
	case valLen >= 1:
		cookieOption := option.Merge(cookieConfig, val[1:])
		c.context.setCookie(
			name,
			val[0].(string),
			cookieOption.Get("maxAge").(int),
			cookieOption.Get("path").(string),
			cookieOption.Get("domain").(string),
			cookieOption.Get("secure").(bool),
			cookieOption.Get("httpOnly").(bool))
	}
	return nil
}
