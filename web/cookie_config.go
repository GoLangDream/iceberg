package web

import . "github.com/GoLangDream/rgo/option"

var cookieConfig = Option{
	"maxAge":   3600,
	"path":     "/",
	"domain":   "localhost",
	"secure":   false,
	"httpOnly": false,
}

// 应该根据配置文件的载入，更新这些参数，特别是 domain
func initCookieConfig() {

}
