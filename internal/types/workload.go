package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// MissionAddonWorkloadEntries 结构体用于 mission_addon_workload_entries 集合
type MissionAddonWorkloadEntries struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Description string             `json:"description" bson:"description"`
	Duration    float64            `json:"duration" bson:"duration"`
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
	UpdatedAt   int64              `json:"updated_at" bson:"updated_at"`
	ProjectID   string             `json:"project_id" bson:"project_id"`
	TaskID      string             `json:"task_id" bson:"task_id"`
	ReportedAt  int64              `json:"reported_at" bson:"reported_at"`
	ProjectInfo *Project           `json:"project_info,omitempty" bson:"project_info,omitempty"`
	TaskInfo    *Task              `json:"task_info,omitempty" bson:"task_info,omitempty"`
}

type WorkloadDTO struct {
	CreatedBy  string `json:"created_by" bson:"created_by"`
	PageSize   int    `json:"page_size" bson:"page_size"`
	PageNumber int    `json:"page_number" bson:"page_number"`
}

type PaginatedWorkload struct {
	Data       []MissionAddonWorkloadEntries `json:"data"`
	Total      int64                         `json:"total"`
	PageSize   int                           `json:"page_size"`
	PageNumber int                           `json:"page_number"`
}
