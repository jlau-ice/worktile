package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// Project 结构体用于 mission_projects 集合
type Project struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}
