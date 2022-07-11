package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type HttpContext struct {
	ctx *fiber.Ctx
}

func (h *HttpContext) text(body string) {
	h.ctx.SendString(body)
}

func (h *HttpContext) json(obj any) {
	h.ctx.JSON(obj)
}

func (h *HttpContext) session() *session.Session {
	session, _ := server.store.Get(h.ctx)
	return session
}

func (h *HttpContext) renderFile(name string, bind any, layouts ...string) {
	h.ctx.Render(name, bind, layouts...)
}

func (h *HttpContext) cookie(name string) string {
	return h.ctx.Cookies(name, "")
}

func (h *HttpContext) setCookie(cookie *fiber.Cookie) {
	h.ctx.Cookie(cookie)
}

func (h *HttpContext) Param(name string, defaultValue ...string) string {
	return h.ctx.Params(name, defaultValue...)
}

func (h *HttpContext) Query(name string, defaultValue ...string) string {
	return h.ctx.Query(name, defaultValue...)
}

func (h *HttpContext) QueryBind(obj any) error {
	return h.ctx.QueryParser(obj)
}

func (h *HttpContext) BodyBind(obj any) error {
	return h.ctx.BodyParser(obj)
}

func (h *HttpContext) FormValue(name string, defaultValue ...string) string {
	return h.ctx.FormValue(name, defaultValue...)
}

func (h *HttpContext) SetHeader(key, value string) {
	h.ctx.Set(key, value)
}

func (h *HttpContext) GetHeader(key string) string {
	return h.ctx.Get(key)
}
