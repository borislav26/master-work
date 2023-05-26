package cv_information

import "cv-maker-service/database"

func (r SimpleRepository) Save(dbManager database.Operations, cvInformation *CVInformation) error {
	return dbManager.Save(cvInformation).Error()
}
