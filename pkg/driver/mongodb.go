package driver

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongodbDriver mongodb 驱动
func MongodbDriver(url string, mongodb string) (*mongo.Database, error) {
	//url = mongodb://localhost:27017
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(mongodb)

	return db, MongodbPing(db)
}

// MongodbPing 测试连接
func MongodbPing(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db.Client().Ping(ctx, readpref.Primary())
}
