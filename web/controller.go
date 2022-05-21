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

	controllers[urlJoin(namespace, name)] = controllerType
}

func doAction(controllerName, actionName string, ctx *HttpContext) {
	controllerType, ok := controllers[controllerName]

	if !ok {
		for cName, cType := range controllers {
			log.Printf("controller [%s], class [%v]", cName, cType)
		}

		log.Printf("调用的controller不存在 %s\n", controllerName)
		return
	}

	controller := reflect.New(controllerType)
	baseController := newBaseController(controllerName, actionName, ctx)
	controller.Elem().FieldByName("BaseController").Set(reflect.ValueOf(baseController))

	action := controller.MethodByName(rstring.Camelize(actionName))
	initController := controller.MethodByName("Init")
	if initController.IsValid() {
		initController.Call(nil)
	}

	if action.IsValid() {
		baseController.beforeActionFilter()
		action.Call([]reflect.Value{})
		baseController.afterActionFilter()
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
