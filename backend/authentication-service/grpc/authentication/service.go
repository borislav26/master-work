package authentication_grpc

import (
	"authentication-service/environement"
	"authentication-service/services/authentication"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (s SimpleService) CheckUserAuthentication(ctx context.Context, request *CheckUserAuthenticationRequest) (*CheckUserAuthenticationResponse, error) {
	token := request.GetUserToken()
	tokenWithoutPrefix := strings.TrimPrefix(token, "Bearer ")
	tokenParts := strings.Split(tokenWithoutPrefix, ".")

	response := &CheckUserAuthenticationResponse{
		Result: false,
	}

	if len(tokenParts) != 3 {
		return response, nil
	}

	header := tokenParts[0]
	payload := tokenParts[1]

	decodedHeaderBytes, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		return response, nil
	}

	var headerMap map[string]any
	err = json.Unmarshal(decodedHeaderBytes, &headerMap)
	if err != nil {
		return response, nil
	}

	if headerMap[authentication.AlgorithmKey] != authentication.HeaderJSON[authentication.AlgorithmKey] {
		return response, nil
	}

	if headerMap[authentication.TypeKey] != authentication.HeaderJSON[authentication.TypeKey] {
		return response, nil
	}

	decodedPayloadBytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return response, nil
	}

	var payloadMap map[string]any
	err = json.Unmarshal(decodedPayloadBytes, &payloadMap)
	if err != nil {
		return response, nil
	}

	userEmail := fmt.Sprintf("%s", payloadMap[authentication.UserEmailKey])

	_, err = s.UserService.FindByEmail(ctx, s.DBManager, userEmail)
	if err != nil {
		return response, err
	}

	expiration := payloadMap[authentication.ExpirationKey]
	expString, ok := expiration.(string)
	if !ok {
		return response, err
	}

	expValue, err := strconv.Atoi(expString)
	if err != nil {
		return response, err
	}

	if int64(expValue) < time.Now().Unix() {
		return response, err
	}

	unsignedHeaderAndPayload := string(decodedHeaderBytes) + string(decodedPayloadBytes)
	hash := hmac.New(sha256.New, []byte(environement.JWT_SECRET_TOKEN_KEY.GetValue()))
	_, err = hash.Write([]byte(unsignedHeaderAndPayload))
	if err != nil {
		return response, nil
	}

	signedHeaderAndPayload := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	if tokenParts[2] != signedHeaderAndPayload {
		return response, nil
	}

	response.Result = true
	return response, nil
}

func (s SimpleService) GetUserEmailFromToken(ctx context.Context, request *CheckUserAuthenticationRequest) (*GetUserEmailFromTokenResponse, error) {
	token := request.UserToken
	tokenWithoutPrefix := strings.TrimPrefix(token, "Bearer ")
	tokenParts := strings.Split(tokenWithoutPrefix, ".")

	response := &GetUserEmailFromTokenResponse{
		Email: "",
	}

	decodedPayload, err := base64.StdEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return nil, errors.New("Failed to decode payload part of token")
	}

	var payloadJSON map[string]string
	err = json.Unmarshal(decodedPayload, &payloadJSON)
	if err != nil {
		return nil, errors.New("Failed to unmarshal payload from token")
	}

	response.Email = payloadJSON[authentication.UserEmailKey]

	return response, nil
}

func (s SimpleService) FindUserByEmail(ctx context.Context, request *FindUserByEmailRequest) (*FindUserByEmailResponse, error) {
	email := request.Email

	response := &FindUserByEmailResponse{}

	user, err := s.UserService.FindByEmail(ctx, s.DBManager, email)
	if err != nil {
		return response, errors.New("Failed to find user by provided email")
	}

	response.Id = int64(user.ID)
	response.Email = user.Email
	response.FirstName = user.FirstName
	response.LastName = user.LastName

	return response, nil
}
