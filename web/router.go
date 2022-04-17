package web

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Router struct {
}

func (r Router) GET(path, to string) {
	toSlice := strings.Split(to, "#")

	server.engine.GET(path, func(ctx *gin.Context) {
		httpContext := HttpContext{ctx}
		doAction(toSlice[0], toSlice[1], &httpContext)
	})
}
