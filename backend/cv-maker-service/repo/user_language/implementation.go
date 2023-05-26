package user_language

import "cv-maker-service/database"

func (r SimpleRepository) Save(dbManager database.Operations, userLanguage *UserLanguage) error {
	return dbManager.Save(userLanguage).Error()
}
