package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	DisplayName    string             `json:"display_name" bson:"display_name"`
	Uid            string             `json:"uid" bson:"uid"`
	Title          string             `json:"title" bson:"title"`
	Gender         string             `json:"gender" bson:"gender"`
	Department     string             `json:"department" bson:"department"`
	DepartmentName string             `json:"department_name,omitempty" bson:"department_name,omitempty"`
	Status         int                `json:"status" bson:"status"`
	IsDeleted      int                `json:"is_deleted" bson:"is_deleted"`
	DisabledAt     int64              `json:"disabled_at" bson:"disabled_at"`
}

type UserPageDTO struct {
	DisplayName string `json:"display_name" form:"display_name"`
	Status      int    `json:"status" form:"status"`
	PageSize    int    `json:"page_size" form:"page_size"`
	PageNumber  int    `json:"page_number" form:"page_number"`
}

type PaginatedUsers struct {
	Data       []User `json:"data"`
	Total      int64  `json:"total"`
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
}
