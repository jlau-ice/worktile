package repository

import (
	"context"
	"worktile/worktile-query-server/internal/types"
	"worktile/worktile-query-server/internal/types/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FetchByName(ctx context.Context, name string) ([]types.User, error) {
	collection := r.db.Collection("users")
	// 还原你原本的正则模糊查询
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
