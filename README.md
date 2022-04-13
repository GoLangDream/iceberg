# iceberg

[![Build](https://github.com/GoLangDream/iceberg/actions/workflows/build.yml/badge.svg)](https://github.com/GoLangDream/iceberg/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/GoLangDream/iceberg/badge.svg?branch=main)](https://coveralls.io/github/GoLangDream/iceberg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoLangDream/iceberg)](https://goreportcard.com/report/github.com/GoLangDream/iceberg)

一个方便快速的web开发框架

## 使用

项目目录

```
├── app.go
├── go.mod
├── go.sum
├── go.work
├── go.work.sum
├── routes
└── web
    ├── controllers
    │   └── home_controller.go
    └── routes.go
```

app.go 内容

```go
package main

import (
	"github.com/GoLangDream/iceberg/web"
	_ "routes/web"
	_ "routes/web/controllers"
)

func main() {
	web.Start()
}
```

Controller 编写
文件 web/controllers/home_controller.go

```go
package controllers

import (
	"github.com/GoLangDream/iceberg/web"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (h *HomeController) Index() {
	h.Text("hello word")
}

func (h *HomeController) SetSession() {
	h.Session("test_session", "a")
}

func (h *HomeController) GetSession() {
	testSession := h.Session("test_session")
	if testSession == nil {
		h.Text("session 没有初始化")
	} else {
		h.Text("session 的值为" + testSession.(string))
	}
}
```

路由编写 
文件 web/routes.go

```go
package web

import (
	. "github.com/GoLangDream/iceberg/web"
)

func init() {
	Routes(func(r Router) {
		r.GET("/hello", "home#index")
		r.GET("/set_session", "home#set_session")
		r.GET("/get_session", "home#get_session")
	})
}
```