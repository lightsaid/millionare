package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/millionare/service/user_srv/model"
)

type userRepo struct {
	col *mongo.Collection
}

func (m *userRepo) Register(user *model.UserModel) (*model.UserModel, error) {
	res, err := m.col.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	uid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}
	user.ID = uid
	return user, nil
}
func (m *userRepo) Login(email, password string, rember bool) (*model.UserModel, error) {
	return nil, nil
}
func (m *userRepo) GetByIDOrEmail(account string) (*model.UserModel, error) {
	return nil, nil
}
