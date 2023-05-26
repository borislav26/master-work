package language

import (
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/repo/language"
)

func (s SimpleService) All(dbManager database.GormDbManager) ([]*language.Language, error) {
	languages, err := s.LanguageRepository.All(dbManager)
	if err != nil {
		return nil, error2.DatabaseError.WrapWithUserMessage(err, "Failed to fetch all languages")
	}
	return languages, nil
}
