package education

import "cv-maker-service/database"

type (
	Education struct {
		database.DatabaseTableRow
		CVDataByteID int64  `column:"fk_cv_data_byte_id" json:"cvDataByteId" gorm:"column:fk_cv_data_byte_id"`
		Faculty      string `column:"faculty" json:"faculty" gorm:"column:faculty"`
		Title        string `column:"title" json:"title" gorm:"column:title"`
		From         string `column:"from" json:"from" gorm:"column:from"`
		To           string `column:"to" json:"to" gorm:"column:to"`
	}

	Repository interface {
		Save(dbManager database.Operations, education *Education) error
	}

	SimpleRepository struct {
	}
)

func (p Education) TableName() string {
	return "educations"
}
