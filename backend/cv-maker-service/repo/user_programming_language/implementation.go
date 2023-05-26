package user_programming_language

import "cv-maker-service/database"

func (r SimpleRepository) Save(dbManager database.Operations, userProgrammingLanguage *UserProgrammingLanguage) error {
	return dbManager.Save(userProgrammingLanguage).Error()
}
