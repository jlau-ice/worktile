// Package handlers backend/handlers/handlers.go
package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	// --- 1. 获取并清洗 IP ---
	// 假设你已经在用 X-Forwarded-For 或是 RealIP 拿到了 "::ffff:192.168.172.92"
	clientIP := c.Request().Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = c.RealIP()
	}
	// 这一步很关键：清洗掉 IPv6 的前缀，只保留 192.168.x.x
	if strings.HasPrefix(clientIP, "::ffff:") {
		clientIP = strings.TrimPrefix(clientIP, "::ffff:")
	}
	// 清洗多重代理的情况 (例如 "192.168.1.1, 10.0.0.1")
	if strings.Contains(clientIP, ",") {
		clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	}
	fmt.Println("最终识别 IP:", clientIP) // 调试用
	// --- 2. 权限“护盾”逻辑 ---
	// 你的 UID
	myHiddenUID := "c1777b3ad3ef4205b3a9c5c043ea6e56"
	// 你的 IP 白名单 (你自己电脑的局域网IP，以及本地回环)
	allowedIPs := map[string]bool{
		"192.168.172.94": true, // 你的局域网 IP
		"127.0.0.1":      true, // 本地调试
		"::1":            true, // 本地调试 IPv6
	}
	// 如果有人试图查询你的 UID
	if uid == myHiddenUID {
		// 检查来源 IP 是否在白名单里
		if !allowedIPs[clientIP] {
			// 只有你能查，别人查直接报错，或者返回空数据装作没这个人
			log.Warn("IP %s 试图偷看你的工时!" + clientIP)
			return c.String(403, "权限不足：无法查看该用户工时")
		}
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
