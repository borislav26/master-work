package cv_data_byte

import (
	"context"
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/repo/cv_data_byte"
	"cv-maker-service/services/authentication"
)

func (s SimpleService) GetAllGeneratedCVsForUser(ctx context.Context, dbManager database.GormDbManager, userEmail string) ([]*cv_data_byte.CVDataByte, error2.Error) {
	request := &authentication.FindUserByEmailRequest{
		Email: userEmail,
	}
	res, err := s.AuthenticationService.FindUserByEmail(ctx, request)
	if err != nil {
		return nil, error2.ServiceLayerError.Wrap(err)
	}

	data, err := s.CVDataByteRepository.AllForUser(dbManager, res.GetId())
	if err != nil {
		return nil, error2.ServiceLayerError.Wrap(err)
	}

	return data, nil
}
