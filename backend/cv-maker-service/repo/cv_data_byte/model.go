package cv_data_byte

import "cv-maker-service/database"

type (
	CVDataByte struct {
		database.DatabaseTableRow
		UserID       int64  `column:"user_id" json:"userId" gorm:"column:user_id"`
		Bytes        []byte `column:"bytes" json:"bytes" gorm:"column:bytes"`
		TemplateName string `column:"template_name" json:"templateName" gorm:"column:template_name"`
	}

	Repository interface {
		Save(dbManager database.Operations, cvDataByte *CVDataByte) error
		AllForUser(operations database.Operations, userID int64) ([]*CVDataByte, error)
	}

	SimpleRepository struct {
	}
)

func (p CVDataByte) TableName() string {
	return "cv_data_bytes"
}
