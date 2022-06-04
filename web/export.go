package web

import "net/http"

// 所有需要导出的内容在这儿定义

type AFH = map[string][]string

var RouterDraw func(router *Router)

func Init() {
	initCookieConfig()
	initServer()
	initRoutes()
}

func Start() {
	server.start()
}

func AllRoutes() []RouterInfo {
	return routes
}

func Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return server.engine.Test(req, msTimeout...)
}
