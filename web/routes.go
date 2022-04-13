package web

func init() {
	// 需要在路由之前注册 session 否则会产生错误
	// 具体参考 https://github.com/gin-contrib/sessions/issues/40
	server.initSession()
}

func Routes(block func(Router)) {
	router := Router{}
	block(router)
}
