package user

import "authentication-service/repo/user"

type AuthenticatedUser struct {
	user.User
	Token string `json:"token"`
}
