package programming_language

import "cv-maker-service/database"

type (
	ProgrammingLanguage struct {
		database.DatabaseTableRow
		Name string `column:"name" json:"name" gorm:"column:name"`
	}

	Repository interface {
		Save(dbManager database.Operations, programmingLanguage *ProgrammingLanguage) error
		FindByName(dbManager database.Operations, name string) (*ProgrammingLanguage, error)
		All(dbManager database.Operations) ([]*ProgrammingLanguage, error)
	}

	SimpleRepository struct {
	}
)

func (p ProgrammingLanguage) TableName() string {
	return "programming_languages"
}
