package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserModel 定义用户模型
type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id" label:"ID"`
	Nickname  string             `bson:"nickname" json:"nickname" label:"昵称"`
	Password  string             `bson:"password" json:"password" label:"密码"`
	Email     string             `bson:"email" json:"email" label:"邮箱"`
	Avatar    string             `bson:"avatar" json:"avatar" label:"头像"`
	Balance   float32            `bson:"balance" json:"balance" label:"余额"`
	Signature string             `bson:"signature" json:"signature" label:"个性签名"` // 个性签名
	Created   time.Time          `bson:"created" json:"created" label:"创建时间"`
	Updated   time.Time          `bson:"updated" json:"updated" label:"更新时间"`
}
