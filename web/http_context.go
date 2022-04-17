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
