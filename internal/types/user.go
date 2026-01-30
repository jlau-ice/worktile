package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// User 结构体用于 users 集合
type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	DisplayName string             `json:"display_name" bson:"display_name"`
	Uid         string             `json:"uid" bson:"uid"`
	// 可以添加其他需要的字段
}
