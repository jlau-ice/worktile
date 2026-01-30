package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 是所有 API 响应的通用结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 返回一个成功的响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

// Error 返回一个错误的响应
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
