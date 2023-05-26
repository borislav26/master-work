package experience

import "cv-maker-service/database"

func (r SimpleRepository) Save(dbManager database.Operations, experience *Experience) error {
	return dbManager.Save(experience).Error()
}
