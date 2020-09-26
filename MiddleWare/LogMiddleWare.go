package MiddleWare

import (
	"bytes"
	"fmt"
	"ginApi/Utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type reqestInfo struct {
	ip     string
	url    string
	method string
	body   string
}

func RequestLog(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	requestInfoObj := reqestInfo{
		c.ClientIP(),
		c.FullPath(),
		c.Request.Method,
		string(body),
	}
	Utils.Log.Info(requestInfoObj)
	//读取出来的数据重新放入body,否则中间件后面的数据读取不到body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
}
