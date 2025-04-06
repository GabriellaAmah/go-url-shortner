package user

import (
	"context"

	"github.com/GabriellaAmah/go-url-shortner/setup"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Email    string `bson:"email,omitempty" json:"email"`
	Password string `bson:"password,omitempty" json:"password"`
	Username string `bson:"username,omitempty" json:"username"`
}

type UserRepository struct { 
	collection *mongo.Collection
}

func (userRepo *UserRepository) initialize(database *mongo.Database) {
	userRepo.collection = database.Collection("user")
}

func (user UserRepository) CreateUser(values User) (*mongo.InsertOneResult, error) {
	data, err := user.collection.InsertOne(context.TODO(), values)
	return data, err
}

func (user UserRepository) FindOne(filter bson.M) (User, error) {
	var result User
	err := user.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}, nil
		}
		return User{}, err
	}
	return result, nil
}

func InitializeUserRepository() UserRepository {
	userRepo := UserRepository{}
	userRepo.initialize(setup.AppConnectionsSetUp())

	return userRepo
}
