package repo

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lightsaid.com/millionare/service/user_srv/model"
)

var collection *mongo.Collection

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("millionare").Collection("user")
}

func TestRegister(t *testing.T) {
	userRepo := NewUserRepo(collection)
	uid := primitive.NewObjectID()
	user := &model.UserModel{
		ID:       uid,
		Nickname: "Lightsaid",
		Password: "abc123",
		Email:    "lightsaid@qq.com",
	}
	res, err := userRepo.Register(user)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, uid, res.ID)
	require.Equal(t, user.Nickname, res.Nickname)
	require.Equal(t, user.Password, res.Password)
	require.Equal(t, user.Email, res.Email)
}

func TestGetEmail(t *testing.T) {
	userRepo := NewUserRepo(collection)
	res, err := userRepo.GetByID("625b1314d1903ca18d326526") //("lightsaid@163.com")
	log.Println(">> ", res, err)
	require.NoError(t, err)

}
