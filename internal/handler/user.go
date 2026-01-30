package handler

import (
	"net/http"
	"worktile/worktile-query-server/internal/types/interfaces"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service interfaces.UserService
}

func NewUserHandel(service interfaces.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUserList(c *gin.Context) {
	ctx := c.Request.Context()
	keyword := c.Query("name")
	users, err := h.service.SearchUsers(ctx, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "查询用户失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": users,
	})
}
