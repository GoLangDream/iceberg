package web

import "github.com/GoLangDream/rgo/rstring"

var routerInfos []RouterInfo

func initRoutes() {
	router := newRootRouter()
	ApplicationRouterDraw(router)
}

func registerRouter(method, path, structName, structMethod, namespace string) {
	routerInfos = append(
		routerInfos,
		RouterInfo{
			"GET",
			path,
			urlJoin(namespace, rstring.Camelize(structName)+"Controller"),
			rstring.Camelize(structMethod),
		})
}

func Routes() []RouterInfo {
	return routerInfos
}
