package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/millionare/service/user_srv/model"
)

// UserRepo 定义 user service dao接口
type UserRepo interface {
	Register(*model.UserModel) (*model.UserModel, error)
	Login(email, password string, rember bool) (*model.UserModel, error)
	GetByIDOrEmail(account string) (*model.UserModel, error)
}

// NewUserRepo 实例化一个 user curd
func NewUserRepo(collection *mongo.Collection) UserRepo {
	return &userRepo{
		col: collection,
	}
}
