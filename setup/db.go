package setup

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabriellaAmah/go-url-shortner/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	database *mongo.Database
}

func (db *DB) ConnectDb() {
	mongodbUrl := config.EnvData.MONGODB_URL
	database := config.EnvData.DATABASE

	if mongodbUrl == "" || database == "" {
		log.Fatal("Set your 'MONGODB_URI' and 'DATABASE' environment variable. ")
	}

	clientOptions := options.Client().ApplyURI(mongodbUrl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB server is not reachable:", err)
	}


	fmt.Println("Mongodb has successfully established connection 🚀",)

	db.database = client.Database(database)
}
