package handler

import (
	"worktile/worktile-query-server/internal/response"
	"worktile/worktile-query-server/internal/types"
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

func (h *UserHandler) GetUsersByPage(c *gin.Context) {
	ctx := c.Request.Context()
	var dto types.UserPageDTO
	if err := c.ShouldBindQuery(&dto); err != nil {
		response.Error(c, 400, "参数绑定失败: "+err.Error())
		return
	}
	result, err := h.service.GetUsersByPage(ctx, dto)
	if err != nil {
		response.Error(c, 500, "分页查询用户失败: "+err.Error())
		return
	}
	response.Success(c, result)
}
