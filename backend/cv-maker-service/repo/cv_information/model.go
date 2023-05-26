package cv_information

import "cv-maker-service/database"

type (
	CVInformation struct {
		database.DatabaseTableRow
		CVDataByteID    int64  `column:"fk_cv_data_byte_id" json:"cvDataByteId" gorm:"column:fk_cv_data_byte_id"`
		JobTitle        string `column:"job_title" json:"jobTitle" gorm:"column:job_title"`
		PhoneNumber     string `column:"phone_number" json:"phoneNumber" gorm:"column:phone_number"`
		Website         string `column:"website" json:"website" gorm:"column:website"`
		LinkedinProfile string `column:"linkedin_profile" json:"linkedinProfile" gorm:"column:linkedin_profile"`
		GithubProfile   string `column:"github_profile" json:"githubProfile" gorm:"column:github_profile"`
		Description     string `column:"description" json:"description" gorm:"description"`
	}

	Repository interface {
		Save(dbManager database.Operations, programmingLanguage *CVInformation) error
	}

	SimpleRepository struct {
	}
)

func (p CVInformation) TableName() string {
	return "cv_information"
}
