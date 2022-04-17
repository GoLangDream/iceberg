package web

import "github.com/gin-gonic/gin"

var server = &app{engine: gin.Default()}

func initServer() {
	server.start()
}
