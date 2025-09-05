package db

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"
	"worktile/worktile-query-server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbURI).SetServerAPIOptions(serverAPI)

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
func GetWorkloadByUserID(userID string) ([]models.WorkloadEntry, error) {
	if !regexp.MustCompile(`^[0-9a-fA-F]{24}$`).MatchString(userID) {
		return nil, fmt.Errorf("无效的用户ID")
	}

	collection := Client.Database(dbName).Collection(workloadCol)
	filter := bson.M{"created_by": userID}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var entries []models.WorkloadEntry
	if err = cursor.All(context.TODO(), &entries); err != nil {
		return nil, err
	}

	// 处理时间戳
	for i := range entries {
		entries[i].CreatedAt = time.Unix(entries[i].CreatedAt, 0).Unix()
		entries[i].UpdatedAt = time.Unix(entries[i].UpdatedAt, 0).Unix()
	}

	return entries, nil
}
