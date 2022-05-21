package web

var loggerFmt = `
Started %s %s %s "%s" for %s at %s
%s
Completed %s %d %s in %s
`

func (s *Server) initLogger() {
	//s.engine.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
	//	var statusColor, methodColor, resetColor string
	//	if params.IsOutputColor() {
	//		statusColor = params.StatusCodeColor()
	//		methodColor = params.MethodColor()
	//		resetColor = params.ResetColor()
	//	}
	//
	//	return fmt.Sprintf(loggerFmt,
	//		methodColor, params.Method, resetColor,
	//		params.Path,
	//		params.ClientIP,
	//		params.TimeStamp.Format(time.RFC1123),
	//		params.ErrorMessage,
	//		statusColor, params.StatusCode, resetColor,
	//		params.Latency,
	//	)
	//}))
}
