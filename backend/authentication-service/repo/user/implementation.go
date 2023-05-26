package user

import (
	"authentication-service/database"
	"database/sql"
)

func (r SimpleRepository) Save(dbManager database.Operations, user *User) error {
	return dbManager.Save(user).Error()
}

func (r SimpleRepository) FindByEmail(dbManager database.Operations, email string) (user *User, err error) {
	err = dbManager.Find(&user, "email = ?", email).Error()
	if user.IsEmpty() {
		err = sql.ErrNoRows
	}
	return
}
