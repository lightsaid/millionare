package models

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TagModel 类型
type TagModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id" label:"ID" binding:"required"`
	Name   string             `bson:"name" json:"name" label:"分类" binding:"required,min=2,max=6"`
	Status string             `bson:"status" json:"status" label:"类型" binding:"required,oneof=Y N"` // Y | N (收入｜支出)
}

// ValidTagStatus 校验 Tag status 是否合法
var ValidTagStatus validator.Func = func(fl validator.FieldLevel) bool {
	if status, ok := fl.Field().Interface().(string); ok {
		return IsSupportType(status)
	}
	return false
}

// IsSupportType tag 类型必须是字符 Y ｜ N
func IsSupportType(status string) bool {
	switch status {
	case "Y", "N":
		return true
	}
	return false
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", ValidTagStatus)
	}
}
