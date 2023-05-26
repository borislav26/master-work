package database

import (
	"time"
)

type (
	Timestamper interface {
		Touch()
	}

	Timestamps struct {
		CreatedAt time.Time `column:"created_at" json:"createdAt"`
		UpdatedAt time.Time `column:"updated_at" json:"updatedAt"`
	}

	Identifier struct {
		ID uint `column:"id" json:"id"`
	}

	GormDbManager interface {
		Operations
		ExecuteTransaction(command Command) error
	}

	DatabaseTableRow struct {
		Identifier
		Timestamps
	}

	Command func(tx Operations) error

	Operations interface {
		TransactionExecutor
	}

	TransactionDBManager struct {
		TDBManager TransactionExecutor
	}
)

func (d DatabaseTableRow) IsEmpty() bool {
	return d.ID == 0
}

func (d DatabaseTableRow) Touch() {
	if d.ID != 0 {
		d.UpdatedAt = time.Now()
	}
}

func (t TransactionDBManager) SelectOne(value any, query string, params ...any) error {
	return t.TDBManager.Select(query, params).Error()
}

func (t *TransactionDBManager) ExecuteTransaction(command Command) error {
	return command(t.TDBManager)
}
