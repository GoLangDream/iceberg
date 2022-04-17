package web

import (
	"fmt"
	. "github.com/GoLangDream/rgo/option"
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

func (c *BaseController) Session(name string, val ...any) any {
	switch len(val) {
	case 0:
		return c.session.Get(name)
	case 1:
		c.session.Set(name, val[0])
		c.session.Save()
	}
	return nil
}

func (c *BaseController) Cookie(name string, val ...any) string {

	switch len(val) {
	case 0:
		return c.context.cookie(name)
	case 1:
		c.setCookie(name, val[0].(string), Option{})
	case 2:
		c.setCookie(name, val[0].(string), val[1].(Option))
	}
	return ""
}

func (c *BaseController) setCookie(name string, value string, option Option) {
	Merge(cookieConfig, option)
	c.context.setCookie(
		name,
		value,
		cookieConfig["maxAge"].(int),
		cookieConfig["path"].(string),
		cookieConfig["domain"].(string),
		cookieConfig["secure"].(bool),
		cookieConfig["httpOnly"].(bool))
}
