package framework

import "github.com/GoLangDream/iceberg/web"

type ApplicationConfig interface {
	RouterDraw() func(*web.Router)
	HomePath() string
}
