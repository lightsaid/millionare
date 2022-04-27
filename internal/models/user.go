package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserModel 定义用户模型
type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id" zhcn:"ID"`
	Nickname  string             `bson:"nickname" json:"nickname" zhcn:"昵称"`
	Password  string             `bson:"password" json:"password" zhcn:"密码"`
	Email     string             `bson:"email" json:"email" zhcn:"邮箱"`
	Avatar    string             `bson:"avatar" json:"avatar" zhcn:"头像"`
	Balance   float32            `bson:"balance" json:"balance" zhcn:"余额"`
	Signature string             `bson:"signature" json:"signature" zhcn:"个性签名"` // 个性签名
	Created   time.Time          `bson:"created" json:"created" zhcn:"创建时间"`
	Updated   time.Time          `bson:"updated" json:"updated" zhcn:"更新时间"`
}
