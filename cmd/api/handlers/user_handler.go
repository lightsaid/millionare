package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lightsaid.com/millionare/internal/models"
	"lightsaid.com/millionare/pkg/errcodes"
)

// RegisterReq 注册入参
type RegisterReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Register 注册请求处理
func (h *APIHandler) Register(c *gin.Context) {
	var req = new(RegisterReq)
	err := c.BindJSON(req)
	if err != nil {
		errcodes.New(
			errcodes.ErrBindBody,
			gin.H{"data": emptyData},
			"请求参数错误", err.Error()).Failed(http.StatusBadRequest, c)
		return
	}

	// 验证必填
	// if msg := user.Require("Email", "Nickname", "Password"); msg != "" {
	// 	response.New(http.StatusBadRequest, nil, msg, nil).Bad(c)
	// 	return
	// }

	// // 验证邮箱
	// if !govalidator.IsEmail(user.Email) {
	// 	response.New(http.StatusBadRequest, nil, "邮箱不正确", nil).Bad(c)
	// 	return
	// }

	// 执行插入mongodb
	user := &models.UserModel{Nickname: req.Nickname, Password: req.Password, Email: req.Email}
	resp, err := h.repo.Register(user)
	if err != nil {
		errcodes.New(
			errcodes.ErrDatabase,
			gin.H{"data": emptyData},
			"请求参数错误", err.Error()).Failed(http.StatusBadRequest, c)
		return
	}
	errcodes.New(
		errcodes.ErrDatabase,
		gin.H{"data": resp},
		"请求参数错误", err.Error()).Success(c)
}

// Login 登录请求处理
// func (h *APIHandler) Login(c *gin.Context) {
// 	user := models.UserModel{
// 		Nickname: "张三",
// 		Email:    "zhansan@qq.com",
// 	}
// 	response.New(http.StatusOK, user, "", nil).OK(c)
// }
