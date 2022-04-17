package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// 需要在路由之前注册 session 否则会产生错误
// 具体参考 https://github.com/gin-contrib/sessions/issues/40
func initSession() {
	store := cookie.NewStore([]byte("secret"))
	server.engine.Use(sessions.Sessions("icebergSession", store))
}
