package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/millionare/internal/models"
)

const dbTimeout = 5 * time.Second

// Repository 所有的 CRUD
type Repository interface {
	// UserModel-------------
	Register(*models.UserModel) (*models.UserModel, error)
	Login(email, password string, rember bool) (*models.UserModel, error)
	GetUserByID(id string) (*models.UserModel, error)
	GetUserByNickname(nickname string) (*models.UserModel, error)
	GetUserByEmail(email string) (*models.UserModel, error)
	UpdateUser(user *models.UserModel) (*models.UserModel, error)
	UpdateUserBalance(id string, balance float32) (*models.UserModel, error)
	UpdateUserAvatar(id string, imgURL string) (*models.UserModel, error)

	// TagModel---------
	InsertTag(*models.TagModel) error
	InsertTagMany([]*models.TagModel) error
	GetTags() ([]models.TagModel, error)
}

type repository struct {
	userRepo
	tagRepo
	mongodb *mongo.Database
}

// NewRepository 创建一个 Repository
func NewRepository(mongodb *mongo.Database) Repository {
	return &repository{
		userRepo: userRepo{
			col: mongodb.Collection("user"),
		},
		tagRepo: tagRepo{
			col: mongodb.Collection("tag"),
		},
	}
}
