package user

import (
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
	rep, err := Services.Grpc.UserGrpc.Login(ctx, &user.LoginReq{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		if rep.Username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "没有找到用户，请先注册"})
			return
		}
	}

	now := time.Now().Unix()

	jwtToken, err := getJwtToken(Services.C.Jwt.Secret, strconv.FormatInt(now, 10), strconv.FormatInt(Services.C.Jwt.Expire, 10), rep.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Name         string
		NickName     string
		AccessToken  string
		AccessExpire int64
		RefreshAfter int64
	}{
		Name:         rep.Username,
		NickName:     rep.Nickname,
		AccessToken:  jwtToken,
		AccessExpire: now + Services.C.Jwt.Expire,
		RefreshAfter: now + Services.C.Jwt.Expire/2})
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
