package web

import "path/filepath"

func urlJoin(paths ...string) string {
	return filepath.Join(paths...)
}

func (r *Router) Namespace(name string, block func(router *Router)) {
	router := &Router{
		urlJoin(r.namespace, name),
		r.scope,
		r.server,
		r.engine.Group(name),
	}

	block(router)
}

func (r *Router) Scope(name string, block func(router *Router)) {
	router := &Router{
		r.namespace,
		urlJoin(r.scope, name),
		r.server,
		r.engine.Group(name),
	}

	block(router)
}
