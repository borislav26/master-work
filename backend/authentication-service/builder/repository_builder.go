package builder

import "authentication-service/repo/user"

type Repositories struct {
	UserRepository user.Repository
}

func BuildRepositories() *Repositories {
	return &Repositories{
		UserRepository: &user.SimpleRepository{},
	}
}
