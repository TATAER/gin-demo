package main

import (
	"ginApi/Dao"
	"ginApi/Router"
	"ginApi/Utils"
)

func main() {
	//日志
	Utils.LogInit()
	Utils.Log.Info("xxxxx")
	//数据库
	Dao.InitMysql()
	defer Dao.Close()
	//路由
	Router.RouterList()
}
