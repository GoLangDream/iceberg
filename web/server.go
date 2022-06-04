package web

import (
	"github.com/GoLangDream/iceberg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/pug"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"net/http"
	"os"
)

type Server struct {
	engine     *fiber.App
	store      *session.Store
	homePath   string
	routerDraw func(router *Router)
	routes     Routes
}

func CreateServer(homePath string, routerDraw func(router *Router)) *Server {
	vConfig, err := viewConfig()
	var engine *fiber.App

	if err == nil {
		vConfig.Debug(false)
		engine = fiber.New(fiber.Config{
			Views:                 vConfig,
			DisableStartupMessage: true,
			ViewsLayout:           "layouts/default",
		})
	} else {
		engine = fiber.New(fiber.Config{
			DisableStartupMessage: true,
			ViewsLayout:           "layouts/default",
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
	log.Info("将启动服务, 监听3000端口, 使用 http://127.0.0.1:3000 访问")
	err := s.engine.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func (s *Server) AllRoutes() []RouterInfo {
	return s.routes.All()
}

func (s *Server) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return s.engine.Test(req, msTimeout...)
}

func viewConfig() (*pug.Engine, error) {
	viewsPath := "web/views"
	dir, err := os.Stat(viewsPath)
	if err != nil || !dir.IsDir() {
		log.Infof("view path %s 不存在", viewsPath)
		return nil, err
	}
	return pug.New(viewsPath, ".pug"), nil

}

func (s *Server) initRoutes() {
	router := newRootRouter(s)
	s.routerDraw(router)
}

func (s *Server) initMiddleware() {
	s.engine.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${url}\n",
	}))
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
	t.SetOutputMirror(os.Stdout)
	t.Style().Format.Header = text.FormatTitle

	t.AppendHeader(table.Row{"#", "Verb", "URI Pattern", "Controller#Action"})

	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:   "Verb",
			Colors: text.Colors{text.BgBlack, text.FgGreen},
		},
	})

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
