package web

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type app struct {
	engine *gin.Engine
}

func (a *app) initSession() {
	fmt.Println("init session")
	store := cookie.NewStore([]byte("secret"))
	a.engine.Use(sessions.Sessions("icebergSession", store))
}

func (a *app) start() {
	a.engine.Run()
}

var server = &app{engine: gin.Default()}

func Start() {
	server.start()
}
