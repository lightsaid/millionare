package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/millionare/internal/repository"
	"lightsaid.com/millionare/pkg/errcodes"
)

// APIHandler 路由 handler
type APIHandler struct {
	repo           repository.Repository
	trans          ut.Translator
	tokenSecretKey string
}

// NewAPIHandler 创建一个APIHandler实例给路由使用
func NewAPIHandler(repo repository.Repository, trans ut.Translator, tokenSecretKey string) *APIHandler {
	return &APIHandler{
		repo:           repo,
		trans:          trans,
		tokenSecretKey: tokenSecretKey,
	}
}

// GetError 从 map[string]string 获取一个err
func (h *APIHandler) GetError(errs map[string]string) error {
	for _, err := range errs {
		if len(err) > 0 {
			return errors.New(err)
		}
	}
	return fmt.Errorf("%s", "未知错误")
}

// BindReqError 处理 gin c.BindXXX 的错误, 如果错误不为nil， 并做出响应
// isWrite 是否已经写入响应，如果写入响应则为 true
func (h *APIHandler) BindReqError(c *gin.Context, err error) (isWrite bool) {
	if err != nil {
		isWrite = true
		errFiled, ok := err.(validator.ValidationErrors)
		if ok {
			err2 := h.GetError(errFiled.Translate(h.trans))
			errcodes.New(
				errcodes.ErrBindBody,
				nil,
				err2.Error(), err.Error()).Failed(http.StatusBadRequest, c)
			return
		}
		errcodes.New(
			errcodes.ErrBindBody,
			nil,
			"请求参数错误", err.Error()).Failed(http.StatusBadRequest, c)
		return
	}
	return
}

// DatabaseError 处理数据库的错误，如果错误不为nil，并做出响应
func (h *APIHandler) DatabaseError(c *gin.Context, err error) {
	if mongo.IsDuplicateKeyError(err) {

	}
}
