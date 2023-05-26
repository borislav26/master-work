package user

import (
	"authentication-service/database"
	error2 "authentication-service/error"
	"authentication-service/repo/user"
	"context"
)

type (
	Service interface {
		Create(ctx context.Context, dbManager database.GormDbManager, email, firstName, lastName, password string) (*AuthenticatedUser, error2.Error)
		LoginUser(ctx context.Context, dbManager database.GormDbManager, email, password string) (*AuthenticatedUser, error2.Error)
		FindByEmail(ctx context.Context, dbManager database.GormDbManager, email string) (*user.User, error2.Error)
	}
	SimpleService struct {
		UserRepository user.Repository
	}
)
