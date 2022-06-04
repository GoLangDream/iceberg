package web

import (
	"fmt"
	"github.com/GoLangDream/iceberg/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var loggerFmt = `
Started %s %s %s "%s" for %s at %s
%s
Completed %s %d %s in %s
`

func (s *Server) initMiddleware() {
	s.engine.Use(recover.New())
	s.engine.Use(logger.New(logger.Config{
		Format: fmt.Sprintf(
			"%s ${ip} ${method} ${url} ${status} ${latency} \n ",
			log.Prefix(),
		),
	}))
}
