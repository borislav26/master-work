package experience

import "cv-maker-service/database"

type (
	Experience struct {
		database.DatabaseTableRow
		Company      string `column:"company" json:"company" gorm:"column:company"`
		JobTitle     string `column:"job_title" json:"jobTitle" gorm:"column:job_title"`
		From         string `column:"from" json:"from" gorm:"column:from"`
		To           string `column:"to" json:"to" gorm:"column:to"`
		Description  string `column:"description" json:"description" gorm:"column:description"`
		CVDataByteID int64  `column:"fk_cv_data_byte_id" json:"cvDataByteId" gorm:"column:fk_cv_data_byte_id"`
	}

	Repository interface {
		Save(dbManager database.Operations, experience *Experience) error
	}

	SimpleRepository struct {
	}
)

func (p Experience) TableName() string {
	return "experiences"
}
