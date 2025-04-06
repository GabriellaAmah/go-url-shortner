package url

import (
	"context"

	"github.com/GabriellaAmah/go-url-shortner/setup"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Url struct {
	Id           string `bson:"_id,omitempty" json:"id"`
	OriginalUrl  string `bson:"originalUrl,omitempty" json:"originalUrl"`
	ShortenedUrl string `bson:"shortenedUrl,omitempty" json:"shortenedUrl"`
	Visits       int    `bson:"visits,omitempty" json:"visits"`
	UserId       string `bson:"userId,omitempty" json:"userId"`
	GuestId      string `bson:"guestId,omitempty" json:"guestId"`
}

type UrlRepository struct {
	collection *mongo.Collection
}

func (urlRepo *UrlRepository) initialize(database *mongo.Database) {
	urlRepo.collection = database.Collection("url")
}

func (urlRepo UrlRepository) CreateUrl(urlData Url) (*mongo.InsertOneResult, error) {
	data, err := urlRepo.collection.InsertOne(context.TODO(), urlData)
	return data, err
}

func (urlRepo UrlRepository) FindOne(filter bson.M) (Url, error) {
	var result Url
	err := urlRepo.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Url{}, nil
		}
		return Url{}, err
	}
	return result, nil
}

func InitializeUrlRepository() UrlRepository {
	urlRepo := UrlRepository{}
	urlRepo.initialize(setup.AppConnectionsSetUp())

	return urlRepo
}
