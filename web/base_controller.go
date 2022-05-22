package web

import (
	"fmt"
	"github.com/GoLangDream/iceberg/database"
	. "github.com/GoLangDream/rgo/option"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
	"path/filepath"
)

type BaseController struct {
	controllerName string
	actionName     string
	context        *HttpContext
	session        *session.Session
	isRender       bool
	beforeActions  []*actionFilter
	afterActions   []*actionFilter
}

func newBaseController(controllerName, actionName string, ctx *HttpContext) *BaseController {
	baseController := BaseController{
		controllerName,
		actionName,
		ctx,
		ctx.session(),
		false,
		[]*actionFilter{},
		[]*actionFilter{},
	}
	return &baseController
}

func (c *BaseController) DB() *gorm.DB {
	return database.DBConn
}

func (c *BaseController) ControllerName() string {
	return c.controllerName
}

func (c *BaseController) ActionName() string {
	return c.actionName
}

func (c *BaseController) Text(body string) {
	c.isRender = true
	if c.context != nil {
		c.context.text(body)
	} else {
		fmt.Println("http context 没有初始化")
	}
}

func (c *BaseController) Json(obj any) {
	c.isRender = true
	c.context.json(obj)
}

func (c *BaseController) render() {
	if !c.isRender && c.context != nil {
		fmt.Printf("render view tmpl file %s", filepath.Join(c.controllerName, c.actionName))
		c.context.renderFile(filepath.Join(c.controllerName, c.actionName), fiber.Map{
			"controller_name": c.controllerName,
			"action_name":     c.actionName,
		})
	}
}

func (c *BaseController) Param(name string) string {
	return c.context.Param(name)
}

func (c *BaseController) Query(name string, defaultValue ...string) string {
	return c.context.Query(name, defaultValue...)
}

func (c *BaseController) FormValue(name string, defaultValue ...string) string {
	return c.context.FormValue(name, defaultValue...)
}

func (c *BaseController) QueryBind(obj any) {
	c.context.QueryBind(obj)
}

func (c *BaseController) BodyBing(obj any) {
	c.context.BodyBind(obj)
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
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Path = cookieConfig["path"].(string)
	cookie.MaxAge = cookieConfig["maxAge"].(int)
	cookie.Domain = cookieConfig["domain"].(string)
	cookie.Secure = cookieConfig["secure"].(bool)
	cookie.HTTPOnly = cookieConfig["httpOnly"].(bool)
	c.context.setCookie(cookie)
}
