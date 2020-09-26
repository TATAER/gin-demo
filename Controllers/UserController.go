package Controllers

import (
	"fmt"
	"ginApi/Entity"
	"ginApi/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
获取用户列表
*/
func UserList(c *gin.Context) {
	userList, err := Entity.GetUserList()
	var data gin.H
	if err != nil {
		data = Utils.ReturnFailed(nil)
	} else {
		for k, v := range userList {
			if v.Id == 8 {
				userList[k].NickName = "xxx"
			}
			userList[k].HeaderImg = ""
		}
		data = Utils.ReturnSuccess(userList)
	}
	c.JSON(http.StatusOK, data)
}

/**
获取用户详情
*/

func UserDetail(c *gin.Context) {
	id := c.Param("id")
	//进行数据类型的转换
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, Utils.ReturnFailed(nil))
	}
	data, err := Entity.GetUserDetailById(userId)
	detail := Entity.WechatUserDetail{
		data,
		"",
	}
	if err != nil {
		c.JSON(http.StatusOK, Utils.ReturnFailed(nil))
	}
	if detail.Id > 2 {
		detail.Level = "高级"
	}
	c.JSON(http.StatusOK, Utils.ReturnSuccess(detail))
}

/**
更新用户信息
*/

func UpdateDetailById(c *gin.Context) {
	//id := c.Param("id")
	var wc Entity.WechatUsers
	c.BindJSON(&wc)
	err := Entity.UpdateDetailById(wc)
	if err != nil {
		c.JSON(http.StatusOK, Utils.ReturnFailed(nil))
	} else {
		c.JSON(http.StatusOK, Utils.ReturnSuccess(nil))
	}
}

/**
个人中心
*/

func Mine(c *gin.Context) {
	userId := c.Param("user_id")
	if userId =="" {
		c.JSON(http.StatusOK, Utils.ReturnFailed(map[string]string{
			"msg": "用户id转换失败",
		}))
		c.Abort()
	}
	fmt.Println(c.Params)
	id,_:=strconv.Atoi(userId)
	wc, _ := Entity.GetUserDetailById(id)
	c.JSON(http.StatusOK, Utils.ReturnSuccess(wc))
}
