package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"worktile/worktile-query-server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client 全局变量，用于在整个应用中共享 MongoDB 客户端
var Client *mongo.Client

// MongoDB连接配置
const (
	// 更新后的数据库连接 URI，密码中的 @ 符号已转义为 %40
	dbURI = "mongodb://mongodb:worktile%40123@192.168.189.207:10000/?authSource=lesschat"
	// 数据库名
	dbName      = "lesschat"
	usersCol    = "users"
	workloadCol = "mission_addon_workload_entries"
)

// InitConnection 初始化 MongoDB 连接
func InitConnection() {
	// 移除 ServerAPI 配置
	opts := options.Client().ApplyURI(dbURI)
	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功连接到 MongoDB!")
	// 调试信息
	db := Client.Database(dbName)
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("获取集合列表失败: %v", err)
	} else {
		fmt.Printf("数据库 %s 中的集合: %v\n", dbName, collections)
	}
}

// GetUsersByName 根据姓名模糊查询用户
func GetUsersByName(name string) ([]models.User, error) {
	collection := Client.Database(dbName).Collection(usersCol)
	filter := bson.M{
		"display_name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}},
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetWorkloadByUserID 根据用户ID查询工时记录
func GetWorkloadByUserID(dto models.WorkloadDTO) ([]models.WorkloadEntry, int64, error) {
	collection := Client.Database(dbName).Collection(workloadCol)
	// 获取总记录数
	total, err := collection.CountDocuments(context.TODO(), bson.M{"created_by": dto.CreatedBy})
	if err != nil {
		return nil, 0, err
	}
	// 构建查询过滤器
	filter := bson.M{"created_by": dto.CreatedBy}
	// 构建分页选项
	skip := int64((dto.PageNumber - 1) * dto.PageSize)
	findOptions := options.Find()
	findOptions.SetLimit(int64(dto.PageSize))
	findOptions.SetSkip(skip)
	// 执行带分页的查询
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.TODO())
	// 将结果解码到实体
	var entries []models.WorkloadEntry
	if err = cursor.All(context.TODO(), &entries); err != nil {
		return nil, 0, err
	}
	return entries, total, nil
}
