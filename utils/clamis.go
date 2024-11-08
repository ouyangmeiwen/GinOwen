package utils

import (
	"net"

	"github.com/gin-gonic/gin"
)

func ClearToken(c *gin.Context) {
	// 增加cookie token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("owentoken", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("owentoken", "", -1, "/", host, false, false)
	}
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("owentoken", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("owentoken", token, maxAge, "/", host, false, false)
	}
}

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("owentoken")
	if token == "" {
		token = c.Request.Header.Get("owentoken")
	}
	return token
}
