package web

import "path/filepath"

func urlJoin(paths ...string) string {
	return filepath.Join(paths...)
}

func (r *Router) Namespace(path string, block func(router Router)) {
	router := Router{
		urlJoin(r.namespace, path),
		r.scope,
		r.engine.Group(path),
	}

	block(router)
}

func (r *Router) Scope(path string, block func(router Router)) {
	router := Router{
		urlJoin(r.namespace, path),
		r.scope,
		r.engine.Group(path),
	}

	block(router)
}
