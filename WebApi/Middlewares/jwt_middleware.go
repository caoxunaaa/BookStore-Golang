package Middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var Secret = "Superxon"

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		mc, ok := ParseToken(authHeader, Secret)
		if ok == false {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		m := mc.(jwt.MapClaims)
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", m["username"])
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func JWTSuperuserMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		mc, ok := ParseToken(authHeader, Secret)
		fmt.Println(mc)
		if ok == false {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		m := mc.(jwt.MapClaims)
		if m["username"] != "管理员" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2006,
				"msg":  "非超级用户",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", m["username"])
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func ParseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}
