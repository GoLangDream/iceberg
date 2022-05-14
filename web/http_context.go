package web

import (
	"github.com/gin-contrib/sessions"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpContext struct {
	c *gin.Context
}

func (h *HttpContext) text(body string) {
	h.c.String(http.StatusOK, body)
}

func (h *HttpContext) session() sessions.Session {
	return sessions.Default(h.c)
}

func (h *HttpContext) cookie(name string) string {
	value, _ := h.c.Cookie(name)
	return value
}

func (h *HttpContext) setCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	h.c.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

func (h *HttpContext) Param(name string) string {
	return h.c.Param(name)
}

func (h *HttpContext) Query(name string) string {
	return h.c.Query(name)
}

func (h *HttpContext) DefaultQuery(name, defaultValue string) string {
	return h.c.DefaultQuery(name, defaultValue)
}

func (h *HttpContext) QueryMap(name string) map[string]string {
	return h.c.QueryMap(name)
}

func (h *HttpContext) PostForm(name string) string {
	return h.c.PostForm(name)
}

func (h *HttpContext) DefaultPostForm(name, defaultValue string) string {
	return h.c.DefaultPostForm(name, defaultValue)
}

func (h *HttpContext) PostFormMap(name string) map[string]string {
	return h.c.PostFormMap(name)
}

func (h *HttpContext) Bind(obj any) {
	h.c.Bind(obj)
}

func (h *HttpContext) ShouldBind(obj any) {
	h.c.ShouldBind(obj)
}

func (h *HttpContext) ShouldBindUri(obj any) {
	h.c.ShouldBindUri(obj)
}

func (h *HttpContext) ShouldBindHeader(obj any) {
	h.c.ShouldBindHeader(obj)
}
