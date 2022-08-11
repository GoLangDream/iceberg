package web

import (
	"github.com/GoLangDream/rgo/rstring"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type Router struct {
	namespace string
	scope     string
	engine    fiber.Router
}

func newRootRouter() *Router {
	return &Router{"", "", server.engine}
}

func (r *Router) handleRequest(httpMethod, path, to string) {
	toSlice := strings.Split(to, "#")

	registerRouter(
		httpMethod,
		urlJoin("/", r.scope, path),
		toSlice[0],
		toSlice[1],
		r.namespace,
	)

	r.engine.Add(httpMethod, path, func(ctx *fiber.Ctx) error {
		httpContext := HttpContext{ctx}
		doAction(
			urlJoin(r.namespace, toSlice[0]),
			toSlice[1],
			&httpContext)
		return nil
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
