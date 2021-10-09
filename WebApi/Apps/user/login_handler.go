package user

import (
	"WebApi/Middlewares"
	"WebApi/Pb/user"
	"WebApi/Services"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func LoginHandler(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	email := c.DefaultPostForm("email", "")
	phone := c.DefaultPostForm("phone", "")

	ctx := context.Background()
	rep, err := Services.UserGrpc.Login(ctx, &user.LoginReq{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		if rep.Ok == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": rep.Code})
			return
		}
	}

	now := time.Now().Unix()
	accessExpire := int64(60 * 60 * 24) // second
	jwtToken, err := getJwtToken(Middlewares.Secret, strconv.FormatInt(now, 10), strconv.FormatInt(accessExpire, 10), username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Name         string
		AccessToken  string
		AccessExpire int64
		RefreshAfter int64
	}{
		Name:         username,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2})
}

func getJwtToken(secretKey string, iat, seconds, userName string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["username"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
