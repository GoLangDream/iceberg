package web

import (
	"strings"

	. "github.com/GoLangDream/iceberg"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r Router) GET(path, to string, options ...Opt) {
	toSlice := strings.Split(to, "#")

	server.engine.GET(path, func(ctx *gin.Context) {
		httpContext := HttpContext{ctx}
		doAction(toSlice[0], toSlice[1], &httpContext)
	})
}
