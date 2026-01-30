package handler

import (
	"worktile/worktile-query-server/internal/response"
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
		response.Error(c, 500, "查询用户失败: "+err.Error())
		return
	}
	response.Success(c, users)
}
