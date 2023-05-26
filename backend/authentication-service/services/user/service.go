package user

import (
	"authentication-service/database"
	error2 "authentication-service/error"
	"authentication-service/repo/user"
	"context"
)

func (s SimpleService) Create(ctx context.Context, dbManager database.GormDbManager, email, firstName, lastName, password string) (*AuthenticatedUser, error2.Error) {
	_, err := s.validateEmail(email)
	if err != nil {
		return nil, error2.DatabaseError.WrapWithUserMessage(err, "Failed to validate email").With("email", email)
	}

	_, err = s.validateNewPassword(password)
	if err != nil {
		return nil, error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to validate password").With("password", password)
	}

	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to hash password").With("password", password)
	}

	newUser := &user.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  hashedPassword,
	}

	err = s.UserRepository.Save(dbManager, newUser)
	if err != nil {
		return nil, error2.DatabaseError.WrapWithUserMessage(err, "Failed to create new user").
			With("email", email).
			With("firstName", firstName).
			With("lastName", lastName)
	}

	if err != nil {
		return nil, error2.DatabaseError.WrapWithUserMessage(err, "Failed to create new user").
			With("email", email).
			With("firstName", firstName).
			With("lastName", lastName)
	}

	user := &AuthenticatedUser{
		User: *newUser,
	}

	err = s.prepareAndSendLog(ctx, "Create", int64(user.ID))
	if err != nil {
		return nil, error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to send log to logger service").With("email", email)
	}

	return user, nil
}

func (s SimpleService) LoginUser(ctx context.Context, dbManager database.GormDbManager, email, password string) (*AuthenticatedUser, error2.Error) {
	user, err := s.UserRepository.FindByEmail(dbManager, email)
	if err != nil {
		return nil, error2.DatabaseError.WrapWithUserMessage(err, "Failed to find user by email = [%s]", email).With("email", email)
	}

	err = s.compareHashAndRealPassword(user.Password, password)
	if err != nil {
		return nil, error2.ServiceLayerError.
			WrapWithUserMessage(err, "Failed to compare real and received password").
			With("realPassword", user.Password).
			With("receivedPassword", password)
	}

	authUser := &AuthenticatedUser{
		User: *user,
	}

	err = s.prepareAndSendLog(ctx, "LoginUser", int64(user.ID))
	if err != nil {
		return nil, error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to send log to logger service").With("email", email)
	}

	return authUser, nil
}

func (s SimpleService) FindByEmail(ctx context.Context, dbManager database.GormDbManager, email string) (*user.User, error2.Error) {
	foundedUser, err := s.UserRepository.FindByEmail(dbManager, email)
	if err != nil {
		return nil, error2.DatabaseError.WrapWithUserMessage(err, "Failed to find user by email").With("email", email)
	}

	err = s.prepareAndSendLog(ctx, "FindByEmail", int64(foundedUser.ID))
	if err != nil {
		return nil, error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to send log to logger service").With("email", email)
	}

	return foundedUser, nil
}
