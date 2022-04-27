package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/millionare/internal/models"
)

// Repository 所有的 CRUD
type Repository interface {
	Register(*models.UserModel) (*models.UserModel, error)
	Login(email, password string, rember bool) (*models.UserModel, error)
	GetUserByID(id string) (*models.UserModel, error)
	GetUserByNickname(nickname string) (*models.UserModel, error)
	GetUserByEmail(email string) (*models.UserModel, error)
}

type repository struct {
	userRepo
	mongodb *mongo.Database
}

// NewRepository 创建一个 Repository
func NewRepository(mongodb *mongo.Database) Repository {
	return &repository{
		userRepo: userRepo{
			col: mongodb.Collection("user"),
		},
	}
}
