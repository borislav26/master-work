package database

import (
	"authentication-service/environement"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBManager struct {
	GORMDBManager   GormDbManager
	CloseConnection func() error
}

func NewDBManager() DBManager {
	dbManager, err := gorm.Open(postgres.Open(environement.DATABASE_URL.GetValue()))
	if err != nil {
		log.Fatalln("Failed to connect to the database")
	}
	sqlDBManager, err := dbManager.DB()
	if err != nil {
		log.Fatalln("Failed to get sql DBManager")
	}
	return DBManager{
		GORMDBManager:   &SimpleTransactionExecutor{dbManager},
		CloseConnection: sqlDBManager.Close,
	}
}

func (dbManager DBManager) Save(value Timestamper) error {
	value.Touch()
	return dbManager.GORMDBManager.Save(value).Error()
}

func (dbManager DBManager) SelectOne(value any, query string, params ...any) error {
	return dbManager.GORMDBManager.Find(value, query, params).Error()
}

func (dbManager DBManager) ExecuteTransaction(command Command) error {
	var err error
	panicHappened := true
	transaction := dbManager.GORMDBManager.Begin()
	if err = transaction.Error(); err != nil {
		return transaction.Error()
	}

	defer func() {
		if err != nil || panicHappened {
			transaction.Rollback()
		}
	}()

	if err = command(dbManager.GORMDBManager); err != nil {
		transaction.Rollback()
	}
	panicHappened = false
	return transaction.Commit().Error()
}
