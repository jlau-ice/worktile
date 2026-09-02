package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"worktile/worktile-query-server/internal/types"
	"worktile/worktile-query-server/internal/types/interfaces"
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FetchByName(ctx context.Context, name string) ([]types.User, error) {
	collection := r.db.Collection("users")
	filter := bson.M{
		"display_name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var users []types.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FetchByPage(ctx context.Context, dto types.UserPageDTO) (*types.PaginatedUsers, error) {
	collection := r.db.Collection("users")

	filter := bson.M{}
	if dto.DisplayName != "" {
		filter["display_name"] = bson.M{"$regex": primitive.Regex{Pattern: dto.DisplayName, Options: "i"}}
	}
	if dto.Status > 0 {
		filter["status"] = dto.Status
	}

	// 获取总数
	total, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	skip := int64((dto.PageNumber - 1) * dto.PageSize)

	// 构建聚合管道以关联部门信息
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$addFields", Value: bson.M{
			"department_oid": bson.M{"$toObjectId": "$department"},
		}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "departments",
			"localField":   "department_oid",
			"foreignField": "_id",
			"as":           "department_info",
		}}},
		{{Key: "$unwind", Value: bson.M{"path": "$department_info", "preserveNullAndEmptyArrays": true}}},
		{{Key: "$addFields", Value: bson.M{
			"department_name": "$department_info.name",
		}}},
		{{Key: "$sort", Value: bson.D{{Key: "disabled_at", Value: -1}, {Key: "_id", Value: 1}}}},
		{{Key: "$skip", Value: skip}},
		{{Key: "$limit", Value: int64(dto.PageSize)}},
		{{Key: "$project", Value: bson.M{
			"department_oid":  0,
			"department_info": 0,
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []types.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return &types.PaginatedUsers{
		Data:       users,
		Total:      total,
		PageSize:   dto.PageSize,
		PageNumber: dto.PageNumber,
	}, nil
}
