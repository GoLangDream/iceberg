package web

import "github.com/GoLangDream/rgo/rstring"

var routerInfos []RouterInfo

func initRoutes() {
	router := newRouter()
	ApplicationRouterDraw(router)
}

func registerRouter(method, path, structName, structMethod string) {
	routerInfos = append(
		routerInfos,
		RouterInfo{
			"GET",
			path,
			rstring.Camelize(structName) + "Controller",
			rstring.Camelize(structMethod),
		})
}

func Routes() []RouterInfo {
	return routerInfos
}
