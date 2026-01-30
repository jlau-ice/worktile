package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task 结构体用于 mission_tasks 集合
type Task struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" bson:"title"`
}
