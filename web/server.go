package web

import (
	"fmt"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/pug"
	"os"
)

var server *webServer

func initServer() {
	// 由于 fiber config 里面的 fiber.Views 是一个 interface类型
	// 但是在判断模板不为空的时候，直接使用的是 Views == nil
	// 就会导致如果手动传入nil进去是不对的
	// 具体可以参考 https://juejin.cn/post/6895231755091968013
	// 于是只好用下面的方法，模拟一个空的interface fiber.Views 传进去
	vConfig := viewConfig()
	var view fiber.Views
	if vConfig != nil {
		view = vConfig
	}

	engine := fiber.New(fiber.Config{
		Views:                 view,
		DisableStartupMessage: true,
		ViewsLayout:           "layouts/default",
	})

	server = &webServer{
		engine: engine,
	}

	server.init()
}

type webServer struct {
	engine *fiber.App
	store  *session.Store
}

func (s *webServer) init() {
	s.initMiddleware()
	s.initSession()
}

func (s *webServer) start() {
	s.printRoutes()
	log.Info("将启动服务, 监听3000端口, 使用 http://127.0.0.1:3000 访问")
	err := s.engine.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func (s *webServer) initSession() {
	s.store = session.New()
}

func (s *webServer) initMiddleware() {
	s.engine.Use(recover.New(recover.Config{
		EnableStackTrace: environment.IsDevelopment(),
	}))
	if !environment.IsTest() {
		s.engine.Use(logger.New(logger.Config{
			Format: fmt.Sprintf(
				"%s ${ip} ${method} ${url} ${status} ${latency} \n ",
				log.Prefix(),
			),
		}))
	}
}

func viewConfig() *pug.Engine {
	viewsPath := "web/views"
	dir, err := os.Stat(viewsPath)
	if err != nil || !dir.IsDir() {
		log.Debugf("view path %s 不存在", viewsPath)
		return nil
	}
	config := pug.New(viewsPath, ".pug")
	//config.Debug(environment.IsDevelopment())
	return config
}
