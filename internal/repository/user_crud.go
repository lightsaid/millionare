package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/millionare/internal/models"
)

type userRepo struct {
	col *mongo.Collection
}

func (m *userRepo) Register(user *models.UserModel) (*models.UserModel, error) {
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
func (m *userRepo) Login(email, password string, rember bool) (*models.UserModel, error) {
	return nil, nil
}
func (m *userRepo) GetUserByID(id string) (*models.UserModel, error) {
	user := new(models.UserModel)
	var err error
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	fileter := bson.M{"_id": oid}
	err = m.col.FindOne(context.Background(), fileter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (m *userRepo) GetUserByNickname(nickname string) (*models.UserModel, error) {
	user := new(models.UserModel)
	var err error
	fileter := bson.M{"nickname": nickname}
	res := m.col.FindOne(context.Background(), fileter)
	err = res.Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (m *userRepo) GetUserByEmail(email string) (*models.UserModel, error) {
	user := new(models.UserModel)
	var err error
	fileter := bson.M{"email": email}
	err = m.col.FindOne(context.Background(), fileter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}
