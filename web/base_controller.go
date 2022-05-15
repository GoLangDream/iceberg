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

func (c *BaseController) Param(name string) string {
	return c.context.Param(name)
}

func (c *BaseController) Query(name string, defaultValue ...string) string {
	switch len(defaultValue) {
	case 0:
		return c.context.Query(name)
	default:
		return c.context.DefaultQuery(name, defaultValue[0])
	}
}

func (c *BaseController) QueryMap(name string) map[string]string {
	return c.context.QueryMap(name)
}

func (c *BaseController) PostForm(name string, defaultValue ...string) string {
	switch len(defaultValue) {
	case 0:
		return c.context.PostForm(name)
	default:
		return c.context.DefaultPostForm(name, defaultValue[0])
	}
}

func (c *BaseController) PostFormMap(name string) map[string]string {
	return c.context.PostFormMap(name)
}

func (c *BaseController) Bind(obj any) {
	c.context.Bind(obj)
}

func (c *BaseController) ShouldBind(obj any) {
	c.context.ShouldBind(obj)
}

func (c *BaseController) ShouldBindUri(obj any) {
	c.context.ShouldBindUri(obj)
}

func (c *BaseController) ShouldBindHeader(obj any) {
	c.context.ShouldBindHeader(obj)
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
