package router

import (
	"worktile/worktile-query-server/internal/handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Params struct {
	dig.In
	UserHandler     *handler.UserHandler
	WorkloadHandler *handler.WorkloadHandler
}

func NewRouter(params Params) *gin.Engine {
	r := gin.New()
	v1 := r.Group("/api")
	{
		RegisterUserRoutes(v1, params.UserHandler)
		RegisterWorkloadRoutes(v1, params.WorkloadHandler)
	}
	return r
}

func RegisterUserRoutes(r *gin.RouterGroup, handler *handler.UserHandler) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", handler.GetUserList)
	}
}

func RegisterWorkloadRoutes(r *gin.RouterGroup, handler *handler.WorkloadHandler) {
	userRoutes := r.Group("/workload")
	{
		userRoutes.GET("/:uid", handler.GetWorkloadList)
	}
}
