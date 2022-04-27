package handlers

import (
	"lightsaid.com/millionare/internal/repository"
)

// APIHandler 路由 handler
type APIHandler struct {
	repo repository.Repository
}

// NewAPIHandler 创建一个APIHandler实例给路由使用
func NewAPIHandler(repo repository.Repository) *APIHandler {
	return &APIHandler{
		repo: repo,
	}
}

var emptyData = struct{}{}
