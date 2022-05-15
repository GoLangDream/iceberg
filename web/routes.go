package web

import (
	"github.com/GoLangDream/rgo/rstring"
)

type Routes struct {
	routerInfos []RouterInfo
}

func (r *Routes) All() []RouterInfo {
	return r.routerInfos
}

func (r *Routes) registerRouter(method, path, structName, structMethod, namespace string) {
	r.routerInfos = append(
		r.routerInfos,
		RouterInfo{
			"GET",
			path,
			urlJoin(namespace, rstring.Camelize(structName)+"Controller"),
			rstring.Camelize(structMethod),
		})
}
