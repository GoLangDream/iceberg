package web

import (
	"github.com/GoLangDream/rgo/rstring"
	"log"
	"reflect"
	"regexp"
	"strings"
)

type Controller interface {
	Text(string)
}

var controllers = make(map[string]reflect.Type)

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
		log.Println("调用的controller不存在")
		return
	}

	controller := reflect.New(controllerType)
	baseController := newBaseController(controllerName, ctx)
	controller.Elem().FieldByName("BaseController").Set(reflect.ValueOf(baseController))

	method := controller.MethodByName(
		rstring.Camelize(actionName))
	if method.IsValid() {
		method.Call([]reflect.Value{})
	} else {
		log.Printf("调用的action [%s]不存在 \n", rstring.Capitalize(actionName))
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
		return rstring.Underscore(strings.Replace(
			typeName, "Controller", "", -1))
	}
	return ""
}
