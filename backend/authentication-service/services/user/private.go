package user

import (
	"authentication-service/environement"
	logger_grpc2 "authentication-service/logs"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"regexp"
)

func (s SimpleService) hashPassword(password string) (string, error) {
	generatedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(generatedPassword), nil
}

func (s SimpleService) validateNewPassword(password string) (match bool, err error) {
	//below pattern matches if password has at least one upper letter, one lower letter and one number
	match, err = regexp.MatchString("^[a-zA-Z0-9_`!@#$%^&.?()-=+]$", password)
	return
}

func (s SimpleService) validateEmail(email string) (match bool, err error) {
	match, err = regexp.MatchString("^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$", email)
	return
}

func (s SimpleService) compareHashAndRealPassword(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return
}

func (s SimpleService) prepareAndSendLog(ctx context.Context, functionName string, userID int64) error {
	request := &logger_grpc2.LogRequest{
		LogEntry: &logger_grpc2.Log{
			ServiceName: environement.SERVICE_NAME.GetValue(),
			Function:    functionName,
			UserId:      userID,
		},
	}
	conn, err := grpc.Dial(environement.GRPC_LOGGER_SERVICE_HOST.GetValue(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}

	loggerClient := logger_grpc2.NewLogServiceClient(conn)
	res, err := loggerClient.WriteLog(ctx, request)
	if err != nil || !res.Result {
		return errors.New("Failed to log action into logger service")
	}

	return nil
}
