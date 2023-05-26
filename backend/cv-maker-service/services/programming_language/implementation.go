package programming_language

import (
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/repo/programming_language"
)

func (s SimpleService) All(dbManager database.GormDbManager) ([]*programming_language.ProgrammingLanguage, error) {
	programmingLanguages, err := s.ProgrammingLanguageRepository.All(dbManager)
	if err != nil {
		return nil, error2.DatabaseError.WrapWithUserMessage(err, "Failed to fetch all programming languages")
	}
	return programmingLanguages, nil
}
