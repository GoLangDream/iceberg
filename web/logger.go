package web

import (
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var loggerFmt = `
Started %s %s %s "%s" for %s at %s
%s
Completed %s %d %s in %s
`

func (s *Server) initLogger() {
	s.engine.Use(recover.New())
}
