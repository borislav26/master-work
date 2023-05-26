package builder

import (
	"authentication-service/services/authentication"
	"authentication-service/services/user"
)

type Services struct {
	UserService           user.Service
	AuthenticationService authentication.Service
}

func BuildServices(r *Repositories) *Services {
	userService := &user.SimpleService{
		UserRepository: r.UserRepository,
	}

	authenticationService := authentication.NewService(userService)

	return &Services{
		AuthenticationService: authenticationService,
		UserService:           userService,
	}
}
