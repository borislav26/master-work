package mongo_repo

import (
	"context"
)

func (r SimpleRepository) SaveLog(context context.Context, record any) error {
	_, err := r.collection.InsertOne(context, record)
	if err != nil {
		return err
	}
	return nil
}
