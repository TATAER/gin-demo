package Router

import (
	"ginApi/Controllers"
	"ginApi/MiddleWare"
	"ginApi/Utils"
	"github.com/gin-gonic/gin"
)

func RouterList() {
	//路由
	r := gin.Default()
	r.Use(MiddleWare.RequestLog)
	usersGroup := r.Group("/user")
	{
		//加入token认证中间件
		usersGroup.Use(Utils.VerifyJwt)
		//获取用户列表
		usersGroup.GET("/list", Controllers.UserList)
		//获取指定用户详情
		usersGroup.GET("/detail/:id", Controllers.UserDetail)
		//获取token对应的个人信息
		usersGroup.GET("/mine", Controllers.Mine)
		//更新用户信息
		usersGroup.POST("/update", Controllers.UpdateDetailById)
	}
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", Controllers.Auth)
	}

	r.Run(":8000")

}
