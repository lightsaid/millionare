package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"lightsaid.com/millionare/internal/models"
	"lightsaid.com/millionare/pkg/errcodes"
	"lightsaid.com/millionare/pkg/security"
)

// RegisterReq 注册入参
type RegisterReq struct {
	Nickname string `json:"nickname" label:"昵称" binding:"required"`
	Email    string `json:"email" label:"邮箱" binding:"required,email"`
	Password string `json:"password" label:"密码" binding:"required,min=6"`
}

// LoginReq 登录入参
type LoginReq struct {
	Email    string `json:"email" label:"邮箱" binding:"required,email"`
	Password string `json:"password" label:"密码" binding:"required,min=6"`
}

// UpdateUserReq 更新用户信息
type UpdateUserReq struct {
	ID        string  `json:"id" label:"id" binding:"required"`
	Avatar    string  `json:"avatar" json:"avatar" label:"头像"`
	Balance   float32 `json:"balance" json:"balance" label:"余额"`
	Signature string  `json:"signature" json:"signature" label:"个性签名"` // 个性签名
}

// UpdateAvatarReq 更新头像
type UpdateAvatarReq struct {
	ID     string `json:"id" label:"id" binding:"required"`
	Avatar string `json:"avatar" json:"avatar" label:"头像" binding:"required"`
}

// Register 注册请求处理
func (h *APIHandler) Register(c *gin.Context) {
	var req = new(RegisterReq)
	err := c.BindJSON(req)
	if ok := h.BindReqError(c, err); ok {
		return
	}
	hashPass, err := security.HashPassword(req.Password)
	if err != nil {
		errcodes.New(errcodes.ErrUnknown, nil, "注册失败，未知错误", err.Error()).Failed(http.StatusBadRequest, c)
		return
	}
	// 添加用户
	user := &models.UserModel{
		Nickname: req.Nickname,
		Password: hashPass,
		Email:    req.Email,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	resp, err := h.repo.Register(user)
	if err != nil {
		var errStr = "注册失败"
		if mongo.IsDuplicateKeyError(err) {
			if strings.Contains(err.Error(), "nickname") {
				errStr = "昵称已经存在"
			}
			if strings.Contains(err.Error(), "email") {
				errStr = "邮箱已经存在"
			}
			errcodes.New(
				errcodes.ErrDatabase,
				nil,
				errStr, err.Error()).Failed(http.StatusBadRequest, c)
			return
		}
		errcodes.New(
			errcodes.ErrDatabase,
			nil,
			errStr, err.Error()).Failed(http.StatusBadRequest, c)
		return
	}
	resp.Password = ""
	errcodes.New(
		errcodes.ErrSuccess,
		gin.H{"data": resp},
		"注册成功", "").Success(c)
}

// Login 登录请求处理
func (h *APIHandler) Login(c *gin.Context) {
	req := new(LoginReq)
	if err := c.BindJSON(req); err != nil {
		if ok := h.BindReqError(c, err); ok {
			return
		}
	}
	// 查询用户是否存在
	user, err := h.repo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errcodes.New(
				errcodes.ErrNotFound,
				gin.H{}, "您还没有注册！",
				err.Error()).Failed(http.StatusNotFound, c)
			return
		}
		errcodes.New(
			errcodes.ErrDatabase,
			gin.H{}, "登录失败，请联系管理员",
			err.Error()).Failed(http.StatusNotFound, c)
		return
	}
	// 验证密码
	err = security.VerifyPassword(user.Password, req.Password)
	user.Password = ""
	if err != nil {
		var errStr = "登录失败，密码无效"
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			errStr = "密码不匹配"
		}
		errcodes.New(
			errcodes.ErrValidation,
			gin.H{}, errStr,
			err.Error()).Failed(http.StatusBadRequest, c)
		return
	}

	// 生成token
	payload, err := security.NewTokenPayload(5*time.Minute, user.Nickname, nil)
	if err != nil {
		errcodes.New(
			errcodes.ErrUnknown,
			gin.H{}, err.Error(),
			err.Error()).Failed(http.StatusBadRequest, c)
		return
	}
	token, err := payload.GenerateToken(h.tokenSecretKey)
	if err != nil {
		errcodes.New(
			errcodes.ErrUnknown,
			gin.H{}, err.Error(),
			err.Error()).Failed(http.StatusBadRequest, c)
		return
	}
	errcodes.New(
		errcodes.ErrSuccess,
		gin.H{"user": user, "token": token},
		"登录成功", "").Success(c)
}

// UpdateUser 更新用户信息
func (h *APIHandler) UpdateUser(c *gin.Context) {
	var req = new(UpdateUserReq)
	if err := c.BindJSON(req); err != nil {
		if ok := h.BindReqError(c, err); ok {
			return
		}
	}
	oid, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		errcodes.New(errcodes.ErrNotFound, nil, "更新失败，ID不合法", err.Error()).Failed(http.StatusBadRequest, c)
		return
	}
	user := &models.UserModel{
		ID:        oid,
		Avatar:    req.Avatar,
		Balance:   req.Balance,
		Signature: req.Signature,
	}
	upUser, err := h.repo.UpdateUser(user)
	if err != nil {
		errcodes.New(errcodes.ErrDatabase, nil, "更新失败", err.Error()).Failed(http.StatusBadRequest, c)
		return
	}
	errcodes.New(errcodes.ErrSuccess, gin.H{"user": upUser}, "更新成功", "").Success(c)
}

// UpdateAvatar 更新头像
func (h *APIHandler) UpdateAvatar(c *gin.Context) {
	var req = new(UpdateAvatarReq)
	if err := c.BindJSON(req); err != nil {
		if ok := h.BindReqError(c, err); ok {
			return
		}
	}
	user, err := h.repo.UpdateUserAvatar(req.ID, req.Avatar)
	if err != nil {
		errcodes.New(errcodes.ErrDatabase, nil, "更新失败", err.Error()).Failed(http.StatusBadRequest, c)
	}
	errcodes.New(errcodes.ErrSuccess, gin.H{"user": user}, "更新成功", "").Success(c)
}
