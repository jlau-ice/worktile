package handler

import (
	"log"
	"net/http"
	"strconv"
	"worktile/worktile-query-server/internal/types"
	"worktile/worktile-query-server/internal/types/interfaces"

	"github.com/gin-gonic/gin"
)

type WorkloadHandler struct {
	service interfaces.WorkloadService
}

func NewWorkloadHandel(service interfaces.WorkloadService) *WorkloadHandler {
	return &WorkloadHandler{
		service: service,
	}
}

func (h *WorkloadHandler) GetWorkloadList(c *gin.Context) {
	ctx := c.Request.Context()
	uid := c.Param("uid")
	clientIP := c.ClientIP()
	myHiddenUID := "c1777b3ad3ef4205b3a9c5c043ea6e56"
	allowedIPs := map[string]bool{
		"192.168.172.94": true,
		"127.0.0.1":      true,
		"::1":            true,
	}
	if uid == myHiddenUID && !allowedIPs[clientIP] {
		log.Printf("非法访问警报: IP %s 试图查看敏感 UID", clientIP)
		c.JSON(http.StatusForbidden, gin.H{"error": "权限不足"})
		return
	}
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNumber, _ = strconv.Atoi(c.DefaultQuery("pageNumber", "1"))

	// 5. 调用 Service
	res, err := h.service.SearchWorkload(ctx, types.WorkloadDTO{
		CreatedBy:  uid,
		PageSize:   pageSize,
		PageNumber: pageNumber,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询负载失败"})
		return
	}
	c.JSON(http.StatusOK, res)
}
