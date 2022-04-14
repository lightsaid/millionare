package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"lightsaid.com/millionare/service/user_srv/repo"
	"lightsaid.com/millionare/service/user_srv/service"
	"lightsaid.com/millionare/service/user_srv/userpb"
)

var address = "0.0.0.0:3333"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	userCol := client.Database("millionare").Collection("user")

	userRepo := repo.NewUserRepo(userCol)
	userService := service.NewUserService(userRepo)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("tcp Listen error: ", err)
	}
	// defer lis.Close()
	srv := grpc.NewServer()
	userpb.RegisterUserServiceServer(srv, userService)

	log.Println("Starting user service on: ", address)
	go func() {
		err = srv.Serve(lis)
		if err != nil {
			log.Println("Starting failed: ", err)
		}
	}()

	// 等到用户终止程序（Ctrl + C)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// 阻塞
	<-ch

	// 停止服务
	srv.Stop()
	// 断开 mongodb 链接，释放资源
	client.Disconnect(context.TODO())
	// 关闭服务监听
	lis.Close()

	log.Println("User Service Stop Successful")
}
