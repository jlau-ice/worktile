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
	// 直接测试 users 集合
	fmt.Printf("正在查询集合: %s.%s\n", dbName, usersCol)
	// 测试集合是否可访问
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		fmt.Printf("统计 users 集合文档失败: %v\n", err)
		return nil, err
	}
	fmt.Printf("users 集合中共有 %d 条记录\n", count)
	var testDoc bson.M
	err = collection.FindOne(context.TODO(), bson.M{}).Decode(&testDoc)
	if err != nil {
		fmt.Printf("从 users 集合获取文档失败: %v\n", err)
		return nil, err
	}
	fmt.Printf("users 集合中的文档示例: %+v\n", testDoc)
	// 现在执行你的查询
	filter := bson.M{
		"display_name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}},
	}
	fmt.Printf("查询过滤器: %+v\n", filter)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("查询失败: %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		fmt.Printf("解码失败: %v\n", err)
		return nil, err
	}
	fmt.Printf("找到 %d 个匹配的用户\n", len(users))
	return users, nil
}

// GetWorkloadByUserID 根据用户ID查询工时记录
// 修改时间戳处理逻辑 08b4cda935554d448bce24f65a5e3a8d
func GetWorkloadByUserID(uid string) ([]models.WorkloadEntry, error) {
	collection := Client.Database(dbName).Collection(workloadCol)
	// 先检查集合中的数据
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("工时记录集合总数: %d\n", count)
	var sampleDoc bson.M
	err = collection.FindOne(context.TODO(), bson.M{}).Decode(&sampleDoc)
	if err != nil {
		fmt.Printf("获取样本文档失败: %v\n", err)
	} else {
		fmt.Printf("样本文档结构: %+v\n", sampleDoc)
	}
	filter := bson.M{"created_by": uid}
	fmt.Printf("查询过滤器: %+v\n", filter)
	var rawResults []bson.M
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &rawResults); err != nil {
		return nil, err
	}
	fmt.Printf("原始查询找到 %d 条记录\n", len(rawResults))
	for i, result := range rawResults {
		fmt.Printf("记录 %d: %+v\n", i+1, result)
	}
	cursor2, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor2.Close(context.TODO())
	var entries []models.WorkloadEntry
	if err = cursor2.All(context.TODO(), &entries); err != nil {
		fmt.Printf("解码到结构体失败: %v\n", err)
		return nil, err
	}
	fmt.Printf("解码后得到 %d 条记录\n", len(entries))
	return entries, nil
}
