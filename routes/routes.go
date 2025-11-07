package routes

import (
	"worktile/worktile-query-server/handlers"

	"github.com/labstack/echo/v4"
)

// InitRoutes 初始化所有路由
func InitRoutes(e *echo.Echo) {
	// 用户管理模块
	userGroup := e.Group("/api/users")
	userGroup.GET("", handlers.SearchUsersHandler)
	// 统计分析模块
	workloadGroup := e.Group("/api/workload")
	workloadGroup.GET("", handlers.GetWorkloadHandler)
}
