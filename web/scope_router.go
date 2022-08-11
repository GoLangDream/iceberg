package web

import "path/filepath"

func urlJoin(paths ...string) string {
	return filepath.Join(paths...)
}

func (r *Router) Namespace(name string, block func(router *Router)) {
	router := &Router{
		urlJoin(r.namespace, name),
		urlJoin(r.scope, r.namespace),
		r.engine.Group(name),
	}

	block(router)
}

func (r *Router) Scope(name string, block func(router *Router)) {
	router := &Router{
		r.namespace,
		urlJoin(r.scope, name),
		r.engine.Group(name),
	}

	block(router)
}
