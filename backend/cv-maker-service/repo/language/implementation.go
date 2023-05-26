package language

import (
	"cv-maker-service/database"
	"database/sql"
)

func (r SimpleRepository) Save(dbManager database.Operations, programmingLanguage *Language) error {
	return dbManager.Save(programmingLanguage).Error()
}

func (r SimpleRepository) FindByName(dbManager database.Operations, name string) (p *Language, err error) {
	err = dbManager.Find(&p, "name = ?", name).Error()
	if p.IsEmpty() {
		err = sql.ErrNoRows
	}
	return
}

func (r SimpleRepository) All(dbManager database.Operations) (languages []*Language, err error) {
	err = dbManager.Find(&languages).Error()
	return
}
