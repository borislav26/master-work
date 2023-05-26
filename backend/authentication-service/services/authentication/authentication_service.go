package authentication

import (
	"authentication-service/database"
	"authentication-service/environement"
	error2 "authentication-service/error"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func (s AuthSimpleService) GenerateToken(context context.Context, dbManager database.GormDbManager, userEmail string) (string, error2.Error) {
	_, err := s.UserService.FindByEmail(context, dbManager, userEmail)
	if err != nil {
		return "", error2.ServiceLayerError.Wrap(err).With("userEmail", userEmail)
	}

	payload := map[string]any{
		"userEmail": userEmail,
		"exp":       fmt.Sprint(time.Now().Add(time.Hour).Unix()),
	}

	return s.GenerateTokenFromPayload(payload)
}

func (s AuthSimpleService) GenerateTokenFromPayload(payload map[string]any) (string, error2.Error) {
	headerBytes, err := json.Marshal(HeaderJSON)
	if err != nil {
		return "", error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to marshal json header")
	}
	headerString := base64.StdEncoding.EncodeToString(headerBytes)

	hash := hmac.New(sha256.New, []byte(environement.JWT_SECRET_TOKEN_KEY.GetValue()))

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to marshal json payload")
	}

	payloadString := base64.StdEncoding.EncodeToString(payloadBytes)
	headerAndPayload := fmt.Sprintf("%s.%s", headerString, payloadString)
	unsignedStr := string(headerBytes) + string(payloadBytes)

	hash.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	tokenStr := fmt.Sprintf("Bearer %s.%s", headerAndPayload, signature)
	return tokenStr, nil
}

func (s AuthSimpleService) FetchUserEmailFromToken(token string) (string, error2.Error) {
	tokenWithoutPrefix := strings.TrimPrefix(token, "Bearer ")
	tokenParts := strings.Split(tokenWithoutPrefix, ".")

	decodedPayload, err := base64.StdEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return "", error2.ServiceLayerError.Wrap(err)
	}

	var payloadJSON map[string]string
	err = json.Unmarshal(decodedPayload, &payloadJSON)
	if err != nil {
		return "", error2.ServiceLayerError.Wrap(err)
	}

	return payloadJSON[UserEmailKey], nil
}

func (s AuthSimpleService) GenerateTokenForUserLogout(context context.Context, dbManager database.GormDbManager, userEmail string) (string, error2.Error) {
	_, err := s.UserService.FindByEmail(context, dbManager, userEmail)
	if err != nil {
		return "", error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to find user by email").With("email", userEmail)
	}

	payload := map[string]any{
		"userEmail": userEmail,
		"exp":       fmt.Sprint(time.Now().Unix()),
	}

	return s.GenerateTokenFromPayload(payload)
}
