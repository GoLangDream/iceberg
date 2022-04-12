package web

import "fmt"

type BaseController struct {
	name    string
	context *HttpContext
}

func (c *BaseController) Text(body string) {
	if c.context != nil {
		c.context.Text(body)
	} else {
		fmt.Println("http context 没有初始化")
	}
}
