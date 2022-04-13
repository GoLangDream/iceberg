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
