package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"log"
)

var server = &app{engine: gin.New()}

func startServer() {
	server.start()
}

func initMiddleware() {
	server.engine.Use(gin.Recovery())
}

func InitServer() {
	initConfig()
	initLogger()
	initMiddleware()
	initSession()
	initRoutes()
}

func printRoutes() {
	t := table.NewWriter()
	t.Style().Format.Header = text.FormatTitle
	t.SetOutputMirror(log.Writer())
	t.AppendHeader(table.Row{"#", "Verb", "URI Pattern", "Controller#Action"})
	for index, info := range Routes() {
		t.AppendRow([]interface{}{
			index + 1,
			info.Method,
			info.Path,
			info.StructName + "#" + info.StructMethod,
		})
	}
	t.Render()
}

func Start() {
	InitServer()
	printRoutes()
	startServer()
}
