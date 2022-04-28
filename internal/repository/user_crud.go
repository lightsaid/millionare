package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/millionare/internal/models"
)

type userRepo struct {
	col *mongo.Collection
}

func (m *userRepo) Register(user *models.UserModel) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	res, err := m.col.InsertOne(ctx, user)
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
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	user := new(models.UserModel)
	var err error
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	fileter := bson.M{"_id": oid}
	err = m.col.FindOne(ctx, fileter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (m *userRepo) GetUserByNickname(nickname string) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	user := new(models.UserModel)
	var err error
	fileter := bson.M{"nickname": nickname}
	res := m.col.FindOne(ctx, fileter)
	err = res.Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (m *userRepo) GetUserByEmail(email string) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	user := new(models.UserModel)
	var err error
	fileter := bson.M{"email": email}
	err = m.col.FindOne(ctx, fileter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (m *userRepo) UpdateUser(user *models.UserModel) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	data, err := m.GetUserByID(user.ID.Hex())
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": data.ID}
	data.Avatar = user.Avatar
	data.Signature = user.Signature
	data.Balance = user.Balance
	data.Updated = time.Now()
	_, err = m.col.ReplaceOne(ctx, filter, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateUserBalance 更新 Balance，参数 balance 可正或负数
func (m *userRepo) UpdateUserBalance(id string, balance float32) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	data, err := m.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	data.Balance += balance

	doc := bson.D{{"$set", bson.D{{"balance", data.Balance}}}}
	_, err = m.col.UpdateOne(ctx, bson.M{"_id": id}, doc)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *userRepo) UpdateUserAvatar(id string, imgURL string) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	data, err := m.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	data.Avatar = imgURL
	doc := bson.D{{"$set", bson.D{{"avatar", data.Avatar}}}}
	_, err = m.col.UpdateOne(ctx, bson.M{"_id": id}, doc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
