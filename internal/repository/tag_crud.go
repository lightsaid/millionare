package repository

import "go.mongodb.org/mongo-driver/mongo"

type tagRepo struct {
	col *mongo.Collection
}
