package programming_language

import (
	"cv-maker-service/database"
	"database/sql"
)

func (r SimpleRepository) Save(dbManager database.Operations, programmingLanguage *ProgrammingLanguage) error {
	return dbManager.Save(programmingLanguage).Error()
}

func (r SimpleRepository) FindByName(dbManager database.Operations, name string) (p *ProgrammingLanguage, err error) {
	err = dbManager.Find(&p, "name = ?", name).Error()
	if p.IsEmpty() {
		err = sql.ErrNoRows
	}
	return
}

func (r SimpleRepository) All(dbManager database.Operations) (languages []*ProgrammingLanguage, err error) {
	err = dbManager.Find(&languages).Error()
	return
}
