package web

import (
	"github.com/GoLangDream/rgo/rstring"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Router struct {
	namespace string
	scope     string
	server    *Server
	engine    gin.IRouter
}

func newRootRouter(server *Server) *Router {
	return &Router{"", "", server, server.engine}
}

func (r *Router) handleRequest(httpMethod, path, to string) {
	toSlice := strings.Split(to, "#")
	if value, ok := r.engine.(*gin.RouterGroup); ok {
		r.server.routes.registerRouter(httpMethod,
			urlJoin(value.BasePath(), path),
			toSlice[0],
			toSlice[1],
			r.namespace,
		)
	} else {
		r.server.routes.registerRouter(
			httpMethod,
			path,
			toSlice[0],
			toSlice[1],
			r.namespace,
		)
	}

	r.engine.Handle(httpMethod, path, func(ctx *gin.Context) {
		httpContext := HttpContext{ctx}
		doAction(
			urlJoin(r.namespace, toSlice[0]),
			toSlice[1],
			&httpContext)
	})
}

func (r *Router) GET(path, to string) {
	r.handleRequest(http.MethodGet, path, to)
}

func (r *Router) PUT(path, to string) {
	r.handleRequest(http.MethodPut, path, to)
}

func (r *Router) POST(path, to string) {
	r.handleRequest(http.MethodPost, path, to)
}

func (r *Router) DELETE(path, to string) {
	r.handleRequest(http.MethodDelete, path, to)
}

func (r *Router) Resources(name string) {
	pluralName := rstring.Plural(name)

	r.GET(pluralName, pluralName+"#index")
	r.GET(pluralName+"/:id", pluralName+"#show")

	r.POST(pluralName, pluralName+"#create")
	r.PUT(pluralName+"/:id", pluralName+"#update")

	r.DELETE(pluralName+"/:id", pluralName+"#destory")
}
