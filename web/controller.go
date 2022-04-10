package web

import (
	"fmt"
	"github.com/GoLangDream/rgo/rstring"
	"reflect"
)

type Controller interface {
	Text(string)
}

var controllers = make(map[string]reflect.Type)

// RegisterController 优化
// 通过 reflect.TypeOf(controller).PkgPath() 获取包名，从而自动注册.
// 这样可以省略第一个 name 参数，根据默认的规则 web/controllers/HomeController 就自动注册成home
func RegisterController(name string, controller Controller) {
	controllerType := reflect.TypeOf(controller)
	controllers[name] = controllerType
}

type BaseController struct {
	name    string
	context *HttpContext
}

func (c *BaseController) Text(body string) {
	if c.context != nil {
		c.context.Text(body)
	} else {
		fmt.Println("http context 没有初始化")
	}
}

func doAction(controllerName string, actionName string, ctx *HttpContext) {
	controllerType, ok := controllers[controllerName]
	controller := reflect.New(controllerType)
	if !ok {
		fmt.Println("调用的controller不存在")
		return
	}

	baseController := &BaseController{controllerName, ctx}
	controller.Elem().FieldByName("BaseController").Set(reflect.ValueOf(baseController))

	method := controller.MethodByName(
		rstring.Capitalize(actionName))
	if method.IsValid() {
		method.Call([]reflect.Value{})
	} else {
		fmt.Printf("调用的action [%s]不存在 \n", rstring.Capitalize(actionName))
	}
}
