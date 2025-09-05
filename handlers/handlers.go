// Package handlers backend/handlers/handlers.go
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"worktile/worktile-query-server/db"
)

// SearchUsersHandler 处理 /api/users 路由
func SearchUsersHandler(c echo.Context) error {
	keyword := c.QueryParam("name")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "name参数不能为空"})
	}

	users, err := db.GetUsersByName(keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "查询用户失败"})
	}

	return c.JSON(http.StatusOK, users)
}

// GetWorkloadHandler 处理 /api/workload/:userID 路由
func GetWorkloadHandler(c echo.Context) error {
	userID := c.Param("userID")

	entries, err := db.GetWorkloadByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, entries)
}
