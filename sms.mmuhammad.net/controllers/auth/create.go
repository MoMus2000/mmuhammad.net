package auth

import "sms.mmuhammad.net/models/model_auth"

type CreateUser struct {
	Email             string
	Password          string
	CreateUserService *model_auth.AuthService
}
