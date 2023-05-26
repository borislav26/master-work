package user_programming_language

import "cv-maker-service/database"

type (
	UserProgrammingLanguage struct {
		database.DatabaseTableRow
		UserID                int64 `column:"user_id" json:"userId" gorm:"column:user_id"`
		CVDataByteID          int64 `column:"fk_cv_data_byte_id" json:"cvDataByteId" gorm:"column:fk_cv_data_byte_id"`
		ProgrammingLanguageID int64 `column:"fk_programming_language_id" json:"programmingLanguageId" gorm:"column:fk_programming_language_id"`
		Percentage            int64 `column:"percentage" json:"percentage" gorm:"column:percentage"`
	}

	Repository interface {
		Save(dbManager database.Operations, programmingLanguage *UserProgrammingLanguage) error
	}

	SimpleRepository struct {
	}
)

func (p UserProgrammingLanguage) TableName() string {
	return "users_programming_languages"
}
