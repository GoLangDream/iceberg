package web

import (
	"github.com/gin-gonic/gin"
)

var server = &app{engine: gin.New()}

func startServer() {
	server.start()
}

func initMiddleware() {
	server.engine.Use(gin.Recovery())
}

func initServer() {
	initConfig()
	initLogger()
	initMiddleware()
	initSession()
	initRoutes()
}

func Start() {
	initServer()
	startServer()
}
