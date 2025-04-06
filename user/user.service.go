package user

import (

	"github.com/GabriellaAmah/go-url-shortner/util"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type IUserService struct {
	userRepo UserRepository
}

func (service IUserService) CreateUser(payload User) (LoginUserResponseI, util.MakeError) {
	emailExists, err := service.userRepo.FindOne(bson.M{"email": payload.Email})
	if err != nil {
		return LoginUserResponseI{}, util.ConstructError(err, "Unable to query user details. please try again", 400)
	}

	// check if username also exists
	if emailExists != (User{}) {
		return LoginUserResponseI{}, util.ConstructError(err, "Email already exists please try another", 400)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return LoginUserResponseI{}, util.ConstructError(err, "An error occurred please try again later", 400)
	}

	password := string(hashedPassword)

	newUser, err := service.userRepo.CreateUser(User{Password: password, Email: payload.Email, Username: payload.Username})
	if err != nil {
		return LoginUserResponseI{}, util.ConstructError(err, "An error occurred please try again later", 400)
	}

	user, err := service.userRepo.FindOne(bson.M{"_id": newUser.InsertedID})
	if err != nil {
		return LoginUserResponseI{}, util.ConstructError(err, "An error occurred please try again later", 400)
	}

	token, err := util.CreateJwtToken(map[string]string{"id": user.Id, "email": user.Email, "username": user.Username})
	if err != nil {
		return LoginUserResponseI{}, util.ConstructError(err, "An error occurred please contact support", 400)
	}

	return LoginUserResponseI{
		Id:    user.Id,
		Email: user.Email,
		Token: token,
	}, util.MakeError{}
}

func (service IUserService) LoginUser(payload LoginUser) (LoginUserResponseI, util.MakeError) {
	emailExists, err := service.userRepo.FindOne(bson.M{"email": payload.Email})
	if err != nil {
		return LoginUserResponseI{}, util.ConstructError(err, "Unable to query user details. please try again", 400)
	}

	if emailExists == (User{}) {
		return LoginUserResponseI{}, util.ConstructError(err, "Invalid email or password", 400)
	}

	hashedPassword := emailExists.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(payload.Password))
	if err != nil {
		return LoginUserResponseI{}, util.ConstructError(err, "Invalid email or password", 400)
	}

	token, err := util.CreateJwtToken(map[string]string{"id": emailExists.Id, "email": emailExists.Email, "username": emailExists.Username})
	if err != nil {
		return LoginUserResponseI{}, util.ConstructError(err, "An error occurred please contact support", 400)
	}

	return LoginUserResponseI{
		Id:    emailExists.Id,
		Email: emailExists.Email,
		Token: token,
	}, util.MakeError{}

}

func (service IUserService) GetUserDetailsById(userId string) (User, util.MakeError) {
	primitiveId, err := util.ConvertStringId(userId)
	if err != nil {
		return User{}, util.ConstructError(err, "Unable to query user details. please try again", 400)
	}

	user, err := service.userRepo.FindOne(bson.M{"_id": primitiveId})
	if err != nil {
		return User{}, util.ConstructError(err, "Unable to query user details. please try again", 400)
	}

	return user, util.MakeError{}
}

var UserService = IUserService{userRepo: InitializeUserRepository()}
