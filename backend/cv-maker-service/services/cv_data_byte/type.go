package cv_data_byte

import (
	"context"
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/repo/cv_data_byte"
	"cv-maker-service/services/authentication"
)

type (
	Service interface {
		GetAllGeneratedCVsForUser(ctx context.Context, dbManager database.GormDbManager, userEmail string) ([]*cv_data_byte.CVDataByte, error2.Error)
	}

	SimpleService struct {
		AuthenticationService authentication.GRPCService
		CVDataByteRepository  cv_data_byte.Repository
	}
)
