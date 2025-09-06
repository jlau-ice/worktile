// Package handlers backend/handlers/handlers.go
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"worktile/worktile-query-server/db"
	"worktile/worktile-query-server/models"
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

// GetWorkloadHandler 处理 /api/workload/:uid 路由
func GetWorkloadHandler(c echo.Context) error {
	uid := c.Param("uid")
	pageSizeStr := c.QueryParam("page_size")
	pageNumberStr := c.QueryParam("page_number")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber <= 0 {
		pageNumber = 1
	}
	// 3. 创建并填充 WorkloadDTO 实体
	dto := models.WorkloadDTO{
		CreatedBy:  uid, // 将 uid 赋值给 CreatedBy 字段
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}
	entries, total, err := db.GetWorkloadByUserID(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	// 构建并返回 PaginatedWorkload 结构体
	response := models.PaginatedWorkload{
		Data:       entries,
		Total:      total,
		PageSize:   dto.PageSize,
		PageNumber: dto.PageNumber,
	}
	return c.JSON(http.StatusOK, response)
}
