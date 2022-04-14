package service

import (
	"context"
	"log"

	"lightsaid.com/millionare/service/user_srv/model"
	"lightsaid.com/millionare/service/user_srv/repo"
	"lightsaid.com/millionare/service/user_srv/userpb"
)

// uerService rpc 实现
type userService struct {
	userpb.UnimplementedUserServiceServer
	userRepo repo.UserRepo
}

// NewUserService 创建一个 userService 给 rpc 调用
func NewUserService(r repo.UserRepo) *userService {
	return &userService{
		userRepo: r,
	}
}

// Register 注册实现
func (m *userService) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	var user = new(model.UserModel)
	user.FromProtoBuff(req.GetUser())

	user, err := m.userRepo.Register(user)
	if err != nil {
		log.Fatal(err)
	}
	resp := user.ToProtoBuff()

	return &userpb.RegisterResponse{User: resp}, nil
}
