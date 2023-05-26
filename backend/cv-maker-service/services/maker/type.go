package maker

import (
	"context"
	"cv-maker-service/database"
	"cv-maker-service/repo/cv_data_byte"
	"cv-maker-service/repo/cv_information"
	"cv-maker-service/repo/education"
	"cv-maker-service/repo/experience"
	"cv-maker-service/repo/language"
	"cv-maker-service/repo/user_language"
	"cv-maker-service/repo/user_programming_language"
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/generator"
	"cv-maker-service/services/mailer"
)

type (
	Service interface {
		CreateCV(ctx context.Context, dbManager database.GormDbManager, cvData *CVData) error
	}

	SimpleService struct {
		AuthenticationService authentication.GRPCService
		GeneratorService      generator.GRPCService
		MailerService         mailer.Service

		CVInformationRepository           cv_information.Repository
		LanguageRepository                language.Repository
		EducationRepository               education.Repository
		UserProgrammingLanguageRepository user_programming_language.Repository
		UserLanguageRepository            user_language.Repository
		ExperienceRepository              experience.Repository
		CVDataByteRepository              cv_data_byte.Repository
	}
)
