package programming_language

import (
	"cv-maker-service/database"
	"cv-maker-service/repo/programming_language"
)

type (
	Service interface {
		All(dbManager database.GormDbManager) ([]*programming_language.ProgrammingLanguage, error)
	}

	SimpleService struct {
		ProgrammingLanguageRepository programming_language.Repository
	}
)
