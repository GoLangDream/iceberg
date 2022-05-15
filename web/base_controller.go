package web

import (
	"fmt"
	. "github.com/GoLangDream/rgo/option"
	"github.com/gin-contrib/sessions"
	"net/http"
	"path/filepath"
)

type BaseController struct {
	controllerName string
	actionName     string
	context        *HttpContext
	session        sessions.Session
	isRender       bool
}

func newBaseController(controllerName, actionName string, ctx *HttpContext) *BaseController {
	baseController := BaseController{
		controllerName,
		actionName,
		ctx,
		ctx.session(),
		false,
	}
	return &baseController
}

func (c *BaseController) Text(body string) {
	if c.context != nil {
		c.isRender = true
		c.context.text(body)
	} else {
		fmt.Println("http context 没有初始化")
	}
}

func (c *BaseController) Json(obj any, code ...int) {
	switch len(code) {
	case 0:
		c.context.json(http.StatusOK, obj)
	default:
		c.context.json(code[0], obj)
	}
}

func (c *BaseController) Render_() {
	if !c.isRender && c.context != nil {
		fmt.Println(
			"call render method controller is " + c.controllerName +
				" action is " + c.actionName +
				"tmpl file is " + filepath.Join(c.controllerName, c.actionName+".html.tmpl"))
		c.context.html(http.StatusOK, filepath.Join(c.controllerName, c.actionName+".html.tmpl"), H{
			"controller_name": c.controllerName,
			"action_name":     c.actionName,
		})
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
