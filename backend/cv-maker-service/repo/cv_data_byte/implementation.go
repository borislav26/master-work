package cv_data_byte

import "cv-maker-service/database"

func (r SimpleRepository) Save(dbManager database.Operations, cvDataByte *CVDataByte) error {
	return dbManager.Save(cvDataByte).Error()
}

func (r SimpleRepository) AllForUser(operations database.Operations, userID int64) (data []*CVDataByte, err error) {
	err = operations.Find(&data, "user_id = ?", userID).Error()
	return
}
