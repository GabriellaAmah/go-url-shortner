package user

type LoginUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponseI struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}