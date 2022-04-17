package web

import . "github.com/GoLangDream/rgo/option"

var cookieConfig = NewOption()

// 应该根据配置文件的载入，更新这些参数，特别是 domain
func initCookieConfig() {
	cookieConfig.Set("maxAge", 3600)
	cookieConfig.Set("path", "/")
	cookieConfig.Set("domain", "localhost")
	cookieConfig.Set("secure", false)
	cookieConfig.Set("httpOnly", false)
}
