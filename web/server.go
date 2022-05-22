package web

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	engine     *fiber.App
	store      *session.Store
	homePath   string
	routerDraw func(router *Router)
	routes     Routes
}

func CreateServer(homePath string, routerDraw func(router *Router)) *Server {
	vConfig, err := viewConfig(homePath)
	vConfig.Debug(true)
	var engine *fiber.App

	if err == nil {
		engine = fiber.New(fiber.Config{
			Views:                 vConfig,
			DisableStartupMessage: true,
			ViewsLayout:           "layouts/main",
		})
	} else {
		engine = fiber.New(fiber.Config{
			DisableStartupMessage: true,
			ViewsLayout:           "layouts/main",
		})
	}

	return &Server{
		engine:     engine,
		homePath:   homePath,
		routerDraw: routerDraw,
	}
}

func (s *Server) InitServer() {
	initCookieConfig()
	s.initLogger()
	s.initMiddleware()
	s.initSession()
	s.initRoutes()
}

func (s *Server) Start() {
	s.InitServer()
	s.printRoutes()
	s.engine.Listen(":3000")
}

func (s *Server) AllRoutes() []RouterInfo {
	return s.routes.All()
}

func (s *Server) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return s.engine.Test(req, msTimeout...)
}

func viewConfig(homePath string) (*html.Engine, error) {
	viewsPath := filepath.Join(homePath, "web/views")
	dir, err := os.Stat(viewsPath)
	if err != nil || !dir.IsDir() {
		fmt.Printf("view path %s 不存在", viewsPath)
		return nil, err
	}
	return html.New(viewsPath, ".html"), nil

}

func (s *Server) initRoutes() {
	router := newRootRouter(s)
	s.routerDraw(router)
}

func (s *Server) initMiddleware() {

}

func (s *Server) initDatabase() {

}

// 需要在路由之前注册 session 否则会产生错误
// 具体参考 https://github.com/gin-contrib/sessions/issues/40
func (s *Server) initSession() {
	s.store = session.New()
}

func (s *Server) printRoutes() {
	t := table.NewWriter()
	t.Style().Format.Header = text.FormatTitle
	t.SetOutputMirror(log.Writer())
	t.AppendHeader(table.Row{"#", "Verb", "URI Pattern", "Controller#Action"})
	for index, info := range s.routes.All() {
		t.AppendRow([]interface{}{
			index + 1,
			info.Method,
			info.Path,
			info.StructName + "#" + info.StructMethod,
		})
	}
	t.Render()
}
