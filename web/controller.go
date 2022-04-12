package web

import (
	"fmt"
	"github.com/GoLangDream/rgo/rstring"
	"reflect"
	"regexp"
	"strings"
)

type Controller interface {
	Text(string)
}

var controllers = make(map[string]reflect.Type)

// RegisterController 优化
// 通过 reflect.TypeOf(controller).PkgPath() 获取包名，从而自动注册.
// 这样可以省略第一个 name 参数，根据默认的规则 web/controllers/HomeController 就自动注册成home
func RegisterController(controller Controller) {
	controllerType := reflect.TypeOf(controller)
	namespace := getNamespace(controllerType.PkgPath())
	name := getName(controllerType.Name())
	if namespace == "" {
		controllers[name] = controllerType
	} else {
		controllers[namespace+"/"+name] = controllerType
	}

}

func doAction(controllerName string, actionName string, ctx *HttpContext) {
	controllerType, ok := controllers[controllerName]

	if !ok {
		fmt.Println("调用的controller不存在")
		return
	}

	controller := reflect.New(controllerType)
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

func getNamespace(packagePath string) string {
	r := regexp.MustCompile(`.*/web/controllers/\/?(.+)\/?`)
	matchArr := r.FindStringSubmatch(packagePath)
	if len(matchArr) >= 2 {
		return strings.TrimRight(matchArr[1], "/")
	}
	return ""
}

func getName(typeName string) string {
	if strings.Contains(typeName, "Controller") {
		return strings.ToLower(strings.Replace(
			typeName, "Controller", "", -1))
	}
	return ""
}
