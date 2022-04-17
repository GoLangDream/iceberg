package web

import . "github.com/GoLangDream/rgo/option"

var cookieConfig = NewOption()

func initCookieConfig() {
	cookieConfig.Set("maxAge", 3600)
	cookieConfig.Set("path", "/")
	cookieConfig.Set("domain", "localhost")
	cookieConfig.Set("secure", false)
	cookieConfig.Set("httpOnly", false)
}
