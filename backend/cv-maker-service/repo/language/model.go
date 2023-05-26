package language

import "cv-maker-service/database"

type (
	Language struct {
		database.DatabaseTableRow
		Name string `column:"name" json:"name" gorm:"column:name"`
	}

	Repository interface {
		Save(dbManager database.Operations, programmingLanguage *Language) error
		FindByName(dbManager database.Operations, name string) (*Language, error)
		All(dbManager database.Operations) ([]*Language, error)
	}

	SimpleRepository struct {
	}
)

func (p Language) TableName() string {
	return "languages"
}
