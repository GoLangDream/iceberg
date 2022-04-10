package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpContext struct {
	c *gin.Context
}

func (h HttpContext) Text(body string) {
	h.c.String(http.StatusOK, body)
}
