package builder

import (
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/cv_data_byte"
	"cv-maker-service/services/generator"
	"cv-maker-service/services/language"
	"cv-maker-service/services/logger"
	"cv-maker-service/services/mailer"
	"cv-maker-service/services/maker"
	"cv-maker-service/services/programming_language"
)

type Services struct {
	GRPCAuthenticationService  authentication.GRPCService
	GRPCGeneratorService       generator.GRPCService
	GRPCLoggerService          logger.GRPCService
	ProgrammingLanguageService programming_language.Service
	LanguageService            language.Service
	MakerService               maker.Service
	MailerService              mailer.GRPCService
	CVDataByteService          cv_data_byte.Service
}

func BuildServices(r *Repositories) *Services {
	programmingLanguageService := &programming_language.SimpleService{
		ProgrammingLanguageRepository: r.ProgrammingLanguageRepository,
	}

	languageService := &language.SimpleService{
		LanguageRepository: r.LanguageRepository,
	}

	grpcAuthenticationService := &authentication.SimpleService{}

	grpcGeneratorService := &generator.SimpleService{}

	grpcLoggerService := &logger.SimpleService{}

	grpcMailerService := &mailer.SimpleService{}

	makerService := &maker.SimpleService{
		AuthenticationService:             grpcAuthenticationService,
		CVInformationRepository:           r.CVInformationRepository,
		LanguageRepository:                r.LanguageRepository,
		UserLanguageRepository:            r.UserLanguageRepository,
		UserProgrammingLanguageRepository: r.UserProgrammingLanguageRepository,
		ExperienceRepository:              r.ExperienceRepository,
		EducationRepository:               r.EducationRepository,
		CVDataByteRepository:              r.CVDataByteRepository,
		GeneratorService:                  grpcGeneratorService,
		MailerService:                     grpcMailerService,
	}

	cvDataByteService := &cv_data_byte.SimpleService{
		CVDataByteRepository:  r.CVDataByteRepository,
		AuthenticationService: grpcAuthenticationService,
	}

	return &Services{
		GRPCAuthenticationService:  grpcAuthenticationService,
		GRPCGeneratorService:       grpcGeneratorService,
		GRPCLoggerService:          grpcLoggerService,
		ProgrammingLanguageService: programmingLanguageService,
		LanguageService:            languageService,
		MakerService:               makerService,
		MailerService:              grpcMailerService,
		CVDataByteService:          cvDataByteService,
	}
}
