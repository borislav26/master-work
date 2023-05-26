package language

import (
	"cv-maker-service/database"
	"cv-maker-service/repo/language"
)

type (
	Service interface {
		All(dbManager database.GormDbManager) ([]*language.Language, error)
	}

	SimpleService struct {
		LanguageRepository language.Repository
	}
)
