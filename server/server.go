package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"worktile/worktile-query-server/db"
)

// NewServer 创建并配置 Echo 实例
func NewServer() *echo.Echo {
	// 初始化数据库连接
	db.InitConnection()
	// 创建 Echo 实例
	e := echo.New()
	// 配置通用中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	return e
}
