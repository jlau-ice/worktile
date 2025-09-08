// Package models
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User 结构体用于 users 集合
type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	DisplayName string             `json:"display_name" bson:"display_name"`
	Uid         string             `json:"uid" bson:"uid"`
	// 可以添加其他需要的字段
}

// WorkloadEntry 结构体用于 mission_addon_workload_entries 集合
type WorkloadEntry struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Description string             `json:"description" bson:"description"`
	Duration    float64            `json:"duration" bson:"duration"`
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
	UpdatedAt   int64              `json:"updated_at" bson:"updated_at"`
	ReportedAt  int64              `json:"reported_at" bson:"reported_at"`
}

type PaginatedWorkload struct {
	Data       []WorkloadEntry `json:"data"`
	Total      int64           `json:"total"`
	PageSize   int             `json:"page_size"`
	PageNumber int             `json:"page_number"`
}

type WorkloadDTO struct {
	CreatedBy  string `json:"created_by" bson:"created_by"`
	PageSize   int    `json:"page_size" bson:"page_size"`
	PageNumber int    `json:"page_number" bson:"page_number"`
}
