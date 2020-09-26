package Utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

var (
	Secret     = "123#111" //salt
	ExpireTime = 3600      //token expire time
)

//generate jwt token
func GenToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New("server busy")
	}
	return signedToken, nil
}

func VerifyJwt(c *gin.Context) {
	var jwtC JWTClaims
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusOK, ReturnFailed(gin.H{
			"msg": "token is empty",
		}))
		c.Abort()
	}
	jwtToken, err := jwt.ParseWithClaims(token, &jwtC, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	claims, ok := jwtToken.Claims.(*JWTClaims)
	if !ok {
		c.JSON(http.StatusOK, ReturnFailed(gin.H{
			"msg": "token 验证失败",
		}))
		c.Abort()
	}
	// 将验证通过的user_id 放到params中
	c.Params = append(c.Params, gin.Param{"user_id", strconv.Itoa(claims.UserId)})
}
