package main

import (
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/millionare/cmd/api/handlers"
	"lightsaid.com/millionare/cmd/api/routes"
	"lightsaid.com/millionare/internal/repository"
	"lightsaid.com/millionare/pkg/driver"
)

var httpAddr = ":4000"
var mongodbURL = "mongodb://localhost:27017"
var mongodbName = "millionare"

func main() {
	var db *mongo.Database
	var err error
	// 1. 链接mongodb
	if db, err = driver.MongodbDriver(mongodbURL, mongodbName); err != nil {
		log.Fatal(err)
	}

	// 2. 创建 Repository
	repo := repository.NewRepository(db)

	// 3. 创建路由handler
	handler := handlers.NewAPIHandler(repo)

	// 4. 创建路由
	r := routes.NewRoutes(handler)

	// 5. web server
	s := &http.Server{
		Addr:           httpAddr,
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 2 * 1024,
	}

	log.Println("Starting HTTP Server on ", httpAddr)

	// 6. 启动
	err = s.ListenAndServe()

	if err != nil {
		log.Println("HTTP Server Start Failed: ", err)
	}
}
