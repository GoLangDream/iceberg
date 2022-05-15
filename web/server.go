package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"log"
	"net/http"
)

type Server struct {
	engine      *gin.Engine
	application Application
	routes      Routes
}

func CreateServer(application Application) *Server {
	return &Server{
		engine:      gin.New(),
		application: application,
	}
}

func (s *Server) InitServer() {
	initConfig()
	s.initLogger()
	s.initMiddleware()
	s.initSession()
	s.initRoutes()
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

func (s *Server) initRoutes() {
	router := newRootRouter(s)
	s.application.RouterDraw(router)
}

func (s *Server) initMiddleware() {
	s.engine.Use(gin.Recovery())
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
