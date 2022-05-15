package web

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	engine     *gin.Engine
	homePath   string
	routerDraw func(router *Router)
	routes     Routes
}

func CreateServer(homePath string, routerDraw func(router *Router)) *Server {
	return &Server{
		engine:     gin.New(),
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
	s.initViews()
}

func (s *Server) Start() {
	s.InitServer()
	s.printRoutes()
	s.engine.Run()
}

func (s *Server) AllRoutes() []RouterInfo {
	return s.routes.All()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.engine.ServeHTTP(w, req)
}

func (s *Server) initViews() {
	viewsPath := filepath.Join(s.homePath, "web/views")
	dir, error := os.Stat(viewsPath)
	if error != nil || !dir.IsDir() {
		return
	}
	viewPath := filepath.Join(
		viewsPath,
		"**/*",
	)
	fmt.Println("views path is " + viewPath)
	s.engine.LoadHTMLGlob(viewPath)
}

func (s *Server) initRoutes() {
	router := newRootRouter(s)
	s.routerDraw(router)
}

func (s *Server) initMiddleware() {
	s.engine.Use(gin.Recovery())
}

func (s *Server) initDatabase() {

}

// 需要在路由之前注册 session 否则会产生错误
// 具体参考 https://github.com/gin-contrib/sessions/issues/40
func (s *Server) initSession() {
	store := cookie.NewStore([]byte("secret"))
	s.engine.Use(sessions.Sessions("icebergSession", store))
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
