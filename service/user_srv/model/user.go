package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightsaid.com/millionare/service/user_srv/userpb"
)

// UserModel 定义用户模型
type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nickname  string             `bson:"nickname" json:"nickname"`
	Password  string             `bson:"password" json:"_"`
	Email     string             `bson:"email" json:"email"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Balance   float32            `bson:"balance" json:"balance"`
	Signature string             `bson:"signature" json:"signature"` // 个性签名
}

// ToProtoBuff usermodel to protobuffer
func (m *UserModel) ToProtoBuff() *userpb.User {
	return &userpb.User{
		Id:        m.ID.Hex(),
		Nickname:  m.Nickname,
		Password:  m.Password,
		Email:     m.Email,
		Avatar:    m.Avatar,
		Balance:   m.Balance,
		Signature: m.Signature,
	}
}

// FromProtoBuff protobuffer to usermodel
func (m *UserModel) FromProtoBuff(user *userpb.User) error {
	uid, err := primitive.ObjectIDFromHex(user.GetId())
	if err != nil {
		return err
	}
	m.ID = uid
	m.Nickname = user.GetNickname()
	m.Password = user.GetPassword()
	m.Email = user.GetEmail()
	m.Avatar = user.GetAvatar()
	m.Balance = user.GetBalance()
	m.Signature = user.GetSignature()
	return nil
}

// CollectionName return mongodb collection name
func CollectionName() string {
	return "user"
}
