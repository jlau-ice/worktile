package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response 是所有 API 响应的通用结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 返回一个成功的响应
func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

// Error 返回一个错误的响应
func Error(c echo.Context, code int, msg string) error {
	return c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// NotFound 返回一个资源未找到的响应
func NotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, Response{
		Code: http.StatusNotFound,
		Msg:  "resource not found",
		Data: nil,
	})
}

// BadRequest 返回一个错误的请求响应
func BadRequest(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, Response{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: nil,
	})
}
