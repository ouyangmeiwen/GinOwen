package utils

import (
	"GINOWEN/global"
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
		c.SetCookie(global.OWEN_CONFIG.System.Token, "", -1, "/", "", false, false)
	} else {
		c.SetCookie(global.OWEN_CONFIG.System.Token, "", -1, "/", host, false, false)
	}
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie(global.OWEN_CONFIG.System.Token, token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie(global.OWEN_CONFIG.System.Token, token, maxAge, "/", host, false, false)
	}
}

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie(global.OWEN_CONFIG.System.Token)
	if token == "" {
		token = c.Request.Header.Get(global.OWEN_CONFIG.System.Token)
	}
	return token
}
