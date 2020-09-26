package Utils

import "github.com/gin-gonic/gin"

/**
统一的成功返回处理
 */
func ReturnSuccess(data interface{}) gin.H {
	return gin.H{
		"code":200,
		"data":data,
	}
}

/**
统一的失败返回处理
 */
func ReturnFailed(data interface{}) gin.H {
	return gin.H{
		"code":100,
		"data":data,
	}
}
