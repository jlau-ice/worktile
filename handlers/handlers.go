// Package handlers backend/handlers/handlers.go
package handlers

import (
	"net/http"
	"strconv"
	"worktile/worktile-query-server/db"
	"worktile/worktile-query-server/models"
	"worktile/worktile-query-server/response"

	"github.com/jlau-ice/gotils/log"
	"github.com/labstack/echo/v4"
)

// SearchUsersHandler 处理 /api/users 路由
func SearchUsersHandler(c echo.Context) error {
	keyword := c.QueryParam("name")
	if keyword == "" {
		return response.Error(c, http.StatusBadRequest, "name参数不能为空")
	}
	log.Info("search users keyword: " + keyword)
	log.Info("search ip: " + c.Request().Host)
	users, err := db.GetUsersByName(keyword)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "查询用户失败")
	}
	return response.Success(c, users)
}

// GetWorkloadHandler 处理 /api/workload/:uid 路由
func GetWorkloadHandler(c echo.Context) error {
	uid := c.QueryParam("uid")
	if uid == "c1777b3ad3ef4205b3a9c5c043ea6e56" {
		return nil
	}
	pageSizeStr := c.QueryParam("pageSize")
	pageNumberStr := c.QueryParam("pageNumber")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber <= 0 {
		pageNumber = 1
	}
	dto := models.WorkloadDTO{
		CreatedBy:  uid,
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}
	entries, total, err := db.GetWorkloadByUserID(dto)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "查询用户工作负载失败")
	}
	paginatedWorkload := models.PaginatedWorkload{
		Data:       entries,
		Total:      total,
		PageSize:   dto.PageSize,
		PageNumber: dto.PageNumber,
	}
	return response.Success(c, paginatedWorkload)
}
