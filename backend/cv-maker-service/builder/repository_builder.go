package builder

import (
	"cv-maker-service/repo/cv_data_byte"
	"cv-maker-service/repo/cv_information"
	"cv-maker-service/repo/education"
	"cv-maker-service/repo/experience"
	"cv-maker-service/repo/language"
	"cv-maker-service/repo/programming_language"
	"cv-maker-service/repo/user_language"
	"cv-maker-service/repo/user_programming_language"
)

type Repositories struct {
	ProgrammingLanguageRepository     programming_language.Repository
	LanguageRepository                language.Repository
	CVInformationRepository           cv_information.Repository
	UserProgrammingLanguageRepository user_programming_language.Repository
	UserLanguageRepository            user_language.Repository
	EducationRepository               education.Repository
	ExperienceRepository              experience.Repository
	CVDataByteRepository              cv_data_byte.Repository
}

func BuildRepositories() *Repositories {
	return &Repositories{
		LanguageRepository:                &language.SimpleRepository{},
		ProgrammingLanguageRepository:     &programming_language.SimpleRepository{},
		CVInformationRepository:           &cv_information.SimpleRepository{},
		UserProgrammingLanguageRepository: &user_programming_language.SimpleRepository{},
		UserLanguageRepository:            &user_language.SimpleRepository{},
		EducationRepository:               &education.SimpleRepository{},
		ExperienceRepository:              &experience.SimpleRepository{},
		CVDataByteRepository:              &cv_data_byte.SimpleRepository{},
	}
}
