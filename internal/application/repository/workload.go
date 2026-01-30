package repository

import (
	"context"
	"worktile/worktile-query-server/internal/types"
	"worktile/worktile-query-server/internal/types/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type workloadRepository struct {
	db *mongo.Database
}

func NewWorkloadRepository(db *mongo.Database) interfaces.WorkloadRepository {
	return &workloadRepository{
		db: db,
	}
}

func (r *workloadRepository) WorkloadByUid(ctx context.Context, dto types.WorkloadDTO) ([]types.MissionAddonWorkloadEntries, int64, error) {
	collection := r.db.Collection("mission_addon_workload_entries")
	filter := bson.M{"created_by": dto.CreatedBy}

	// 1. 获取总数
	total, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	// 2. 分页偏移量
	skip := int64((dto.PageNumber - 1) * dto.PageSize)
	// 3. 构建聚合管道 (这里直接复用你原来的逻辑，但注意 context 传递)
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$lookup", Value: bson.M{
			"from": "mission_projects", "localField": "project_id", "foreignField": "_id", "as": "project_info",
		}}},
		{{Key: "$lookup", Value: bson.M{
			"from": "mission_tasks", "localField": "task_id", "foreignField": "_id", "as": "task_info",
		}}},
		{{Key: "$unwind", Value: bson.M{"path": "$project_info", "preserveNullAndEmptyArrays": true}}},
		{{Key: "$unwind", Value: bson.M{"path": "$task_info", "preserveNullAndEmptyArrays": true}}},
		{{Key: "$addFields", Value: bson.M{
			"project_name": "$project_info.name",
			"task_title":   "$task_info.title",
		}}},
		{{Key: "$sort", Value: bson.D{{Key: "reported_at", Value: -1}}}},
		{{Key: "$skip", Value: skip}},
		{{Key: "$limit", Value: int64(dto.PageSize)}},
	}

	// 4. 执行聚合
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var entries []types.MissionAddonWorkloadEntries
	if err := cursor.All(ctx, &entries); err != nil {
		return nil, 0, err
	}
	return entries, total, nil
}
