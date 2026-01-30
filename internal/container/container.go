package container

import (
	"context"
	"fmt"
	"time"
	"worktile/worktile-query-server/internal/application/repository"
	"worktile/worktile-query-server/internal/application/service"
	"worktile/worktile-query-server/internal/config"
	"worktile/worktile-query-server/internal/handler"
	"worktile/worktile-query-server/internal/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

func BuildContainer(container *dig.Container) *dig.Container {
	must(container.Provide(config.LoadConfig))
	must(container.Provide(initDatabase))
	must(container.Provide(func(client *mongo.Client, cfg *config.Config) *mongo.Database {
		return client.Database(cfg.Database.DBName)
	}))
	must(container.Provide(handler.NewUserHandel))
	must(container.Provide(handler.NewWorkloadHandel))
	must(container.Provide(service.NewUserService))
	must(container.Provide(service.NewWorkloadService))
	must(container.Provide(repository.NewWorkloadRepository))
	must(container.Provide(repository.NewUserRepository))
	must(container.Provide(router.NewRouter))
	return container
}

func initDatabase(cfg *config.Config) (*mongo.Client, error) {
	dsn := cfg.GetDSN()
	// 2. 配置 MongoDB 客户端选项
	clientOptions := options.Client().ApplyURI(dsn)
	// 3. 设置连接超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 4. 建立连接
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("无法连接到 MongoDB: %w", err)
	}
	// 5. 检查连接是否可用
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("MongoDB Ping 失败: %w", err)
	}
	return client, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
