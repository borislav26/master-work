package user_language

import "cv-maker-service/database"

type (
	UserLanguage struct {
		database.DatabaseTableRow
		UserID       int64 `column:"user_id" json:"userId" gorm:"column:user_id"`
		CVDataByteID int64 `column:"fk_cv_data_byte_id" json:"cvDataByteId" gorm:"column:fk_cv_data_byte_id"`
		LanguageID   int64 `column:"fk_language_id" json:"languageId" gorm:"column:fk_language_id"`
	}

	Repository interface {
		Save(dbManager database.Operations, userLanguage *UserLanguage) error
	}

	SimpleRepository struct {
	}
)

func (p UserLanguage) TableName() string {
	return "users_languages"
}
