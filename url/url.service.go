package url

import (
	"math/rand"

	"github.com/GabriellaAmah/go-url-shortner/config"
	"github.com/GabriellaAmah/go-url-shortner/user"
	"github.com/GabriellaAmah/go-url-shortner/util"
	"go.mongodb.org/mongo-driver/bson"
)


type IUrlService struct {
	urlRepo UrlRepository
	userService user.IUserService
}

func (service IUrlService) CreateShortenUrl(userId string, payload IShortenUrl) (Url, util.MakeError) {
	
	userExists, error := service.userService.GetUserDetailsById(userId)
	if error.IsError {
		return Url{}, util.ConstructError(error, "An error occurred please try again", 400)
	}

	if userExists == (user.User{}){
		return Url{}, util.ConstructError(error, "User data does not exists", 400)
	}

	shortenedUrl := service.shortenUrl()

	 savedUrl, err :=service.urlRepo.CreateUrl(Url{
		OriginalUrl: payload.Url,
		ShortenedUrl: config.EnvData.BASE_URL + shortenedUrl,
		Visits: 0,
		UserId: userExists.Id,
	})
	if err != nil {
		return Url{}, util.ConstructError(error, "An error occurred while saving data", 400)
	}

	data, err := service.urlRepo.FindOne(bson.M{"_id": savedUrl.InsertedID})
	if err != nil {
		return Url{}, util.ConstructError(error, "An error occurred while reading data", 400)
	}

	return data, util.MakeError{}

}

func (service IUrlService) shortenUrl() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    const keyLength = 6

    shortKey := make([]byte, keyLength)
    for i := range shortKey {
        shortKey[i] = charset[rand.Intn(len(charset))]
    }
    return string(shortKey)
}

var UrlService = IUrlService{urlRepo: InitializeUrlRepository(), userService: user.UserService}
