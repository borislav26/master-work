package user

import (
	"authentication-service/database"
)

type (
	User struct {
		database.DatabaseTableRow
		Email     string `column:"email" json:"email" gorm:"column:email"`
		FirstName string `column:"first_name" json:"firstName" gorm:"column:first_name"`
		LastName  string `column:"last_name" json:"lastName" gorm:"column:last_name"`
		Password  string `column:"password" json:"password" gorm:"column:password"`
	}

	Repository interface {
		Save(dbManager database.Operations, user *User) error
		FindByEmail(dbManager database.Operations, email string) (*User, error)
	}

	SimpleRepository struct {
	}
)

func (u User) TableName() string {
	return "users"
}
