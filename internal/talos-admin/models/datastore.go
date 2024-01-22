package models

import (
	"context"

	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/database"
	"github.com/uptrace/bun"
)

type DataStore struct {
	db  *bun.DB
	ctx context.Context
}

func NewDatastore(cfg config.Сonfig) (DataStore, error) {
	var ds DataStore

	ds.db = database.New(cfg)
	ds.ctx = context.TODO()

	return ds, nil
}

func (ds *DataStore) Init() error {
	_, err := ds.db.NewCreateTable().Model((*Cluster)(nil)).Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}
