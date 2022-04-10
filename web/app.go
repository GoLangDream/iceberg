package web

import "github.com/gin-gonic/gin"

type app struct {
	engine *gin.Engine
}

func (a app) start() {
	a.engine.Run()
}

var server = app{engine: gin.Default()}

func Start() {
	server.start()
}
