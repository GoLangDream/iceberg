package web

import (
	"github.com/GoLangDream/rgo/rstring"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

var routes []RouterInfo

func initRoutes() {
	if RouterDraw != nil {
		RouterDraw(newRootRouter())
	}
}

func registerRouter(method, path, structName, structMethod, namespace string) {
	routes = append(
		routes,
		RouterInfo{
			"GET",
			path,
			urlJoin(namespace, rstring.Camelize(structName)+"Controller"),
			rstring.Camelize(structMethod),
		})
}

func (s *webServer) printRoutes() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.Style().Format.Header = text.FormatTitle

	t.SetTitle("网站路由")
	t.Style().Title.Align = text.AlignCenter
	t.AppendHeader(table.Row{"#", "Verb", "URI Pattern", "Controller#Action"})

	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:   "Verb",
			Colors: text.Colors{text.BgBlack, text.FgGreen},
		},
	})

	for index, info := range routes {
		t.AppendRow([]interface{}{
			index + 1,
			info.Method,
			info.Path,
			info.StructName + "#" + info.StructMethod,
		})
	}
	t.Render()
}
