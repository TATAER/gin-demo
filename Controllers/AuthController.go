package Controllers

import (
	"ginApi/Entity"
	"ginApi/Modle"
	"ginApi/Utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	var loginModel Modle.Login
	c.BindJSON(&loginModel)
	where := map[string]interface{}{
		"nick_name": loginModel.UserName,
		"phone":     loginModel.PassWord,
	}
	if id, ok := Entity.Auth(where); ok {
		sc := jwt.StandardClaims{
			ExpiresAt: int64(Utils.ExpireTime),
		}
		jwtClaims := Utils.JWTClaims{
			sc,
			id,
		}
		jwtToken, err := Utils.GenToken(&jwtClaims)
		if err != nil {
			c.JSON(http.StatusOK, Utils.ReturnFailed(map[string]string{
				"msg": "token生成失败",
			}))
		} else {
			c.JSON(http.StatusOK, Utils.ReturnSuccess(map[string]string{
				"jwtToken": jwtToken,
			}))
		}
	} else {
		c.JSON(http.StatusOK, Utils.ReturnFailed(map[string]string{
			"msg": "登陆失败",
		}))
	}

}
