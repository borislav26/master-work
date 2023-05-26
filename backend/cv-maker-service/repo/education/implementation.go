package education

import "cv-maker-service/database"

func (r SimpleRepository) Save(dbManager database.Operations, education *Education) error {
	return dbManager.Save(education).Error()
}
