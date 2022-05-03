package web

import (
	"github.com/GoLangDream/rgo/rstring"
	"github.com/gin-gonic/gin"
	"strings"
)

type Router struct {
}

func newRouter() Router {
	return Router{}
}

func (r *Router) GET(path, to string) {
	toSlice := strings.Split(to, "#")
	registerRouter("GET", path, toSlice[0], toSlice[1])

	server.engine.GET(path, func(ctx *gin.Context) {
		httpContext := HttpContext{ctx}
		doAction(toSlice[0], toSlice[1], &httpContext)
	})
}

func (r *Router) PUT(path, to string) {
	toSlice := strings.Split(to, "#")
	registerRouter("PUT", path, toSlice[0], toSlice[1])

	server.engine.PUT(path, func(ctx *gin.Context) {
		httpContext := HttpContext{ctx}
		doAction(toSlice[0], toSlice[1], &httpContext)
	})
}

func (r *Router) POST(path, to string) {
	toSlice := strings.Split(to, "#")

	registerRouter("POST", path, toSlice[0], toSlice[1])
	server.engine.POST(path, func(ctx *gin.Context) {
		httpContext := HttpContext{ctx}
		doAction(toSlice[0], toSlice[1], &httpContext)
	})
}

func (r *Router) DELETE(path, to string) {
	toSlice := strings.Split(to, "#")

	registerRouter("DELETE", path, toSlice[0], toSlice[1])
	server.engine.DELETE(path, func(ctx *gin.Context) {
		httpContext := HttpContext{ctx}
		doAction(toSlice[0], toSlice[1], &httpContext)
	})
}

func (r *Router) Resources(name string) {
	pluralName := rstring.Plural(name)
	//singularName := rstring.Singular(name)

	r.GET(pluralName, pluralName+"#index")
	r.GET(pluralName+"/:id", pluralName+"#show")

	r.POST(pluralName, pluralName+"#create")
	r.PUT(pluralName+"/:id", pluralName+"#update")

	r.DELETE(pluralName+"/:id", pluralName+"#destory")

}
