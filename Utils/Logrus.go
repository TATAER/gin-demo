package Utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log *logrus.Logger

func LogInit() {
	Log = logrus.New()
	Log.Formatter = &logrus.JSONFormatter{}
	date := time.Now().Format("20060102")
	logFile := "gin" + date + ".text"
	f, _ := os.Create(logFile)
	gin.SetMode(gin.ReleaseMode)
	//框架请求日志记录到文件
	Log.Out = f
	gin.DefaultWriter = Log.Out
	// 日志级别为info
	Log.Level = logrus.InfoLevel
}
