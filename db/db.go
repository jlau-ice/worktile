package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config 数据库连接配置
type Config struct {
	URI    string
	DBName string
}

// Connect 创建并返回一个新的 MongoDB 客户端
func Connect(ctx context.Context, cfg Config) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	// Ping 数据库以验证连接
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		err := client.Disconnect(context.Background())
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	log.Println("成功连接到 MongoDB!")
	return client, nil
}
