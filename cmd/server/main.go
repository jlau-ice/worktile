package main

import (
	"context"
	"log"
	"time"
	"worktile/worktile-query-server/internal/config"
	"worktile/worktile-query-server/internal/container"
	"worktile/worktile-query-server/internal/runtime"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	c := container.BuildContainer(runtime.GetContainer())
	err := c.Invoke(func(
		cfg *config.Config,
		db *mongo.Client,
		router *gin.Engine,
	) error {
		// 确保数据库连接在程序结束时关闭
		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := db.Disconnect(ctx); err != nil {
				log.Printf("关闭 MongoDB 连接失败: %v", err)
			} else {
				log.Println("MongoDB 连接已成功关闭")
			}
		}()
		// 启动服务器
		port := ":" + cfg.Server.Port
		log.Printf("服务器启动在端口 %s", port)
		return router.Run(port)
	})
	if err != nil {
		log.Fatalf("应用启动失败: %v", err)
	}
}
