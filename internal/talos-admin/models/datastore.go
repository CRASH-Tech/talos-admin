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
	_, err := ds.db.NewCreateTable().IfNotExists().Model((*Cluster)(nil)).Exec(ds.ctx)
	if err != nil {
		return err
	}

	_, err = ds.db.NewCreateTable().IfNotExists().Model((*Node)(nil)).Exec(ds.ctx)
	if err != nil {
		return err
	}

	_, err = ds.db.NewCreateTable().IfNotExists().Model((*Template)(nil)).Exec(ds.ctx)
	if err != nil {
		return err
	}

	_, err = ds.db.NewCreateTable().IfNotExists().Model((*Variable)(nil)).Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DataStore) ReadAll(target interface{}, result interface{}, limit, offset int, sort string) (int, error) {
	q := ds.db.NewSelect().Model(target)

	if limit >= 0 {
		q = q.Limit(limit)
	}

	if offset >= 0 {
		q = q.Offset(offset)
	}

	if sort != "" {
		q = q.OrderExpr(sort)
	}

	count, err := q.ScanAndCount(ds.ctx, result)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ds *DataStore) Read(id int64, target interface{}) error {
	err := ds.db.NewSelect().
		Model(target).
		Where("id = ?", id).
		Scan(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DataStore) Create(target interface{}) (int64, error) {
	var id int64
	_, err := ds.db.NewInsert().Model(target).Returning("id").Exec(ds.ctx, &id)
	if err != nil {
		return id, err
	}

	return id, nil

}

func (ds *DataStore) Delete(id int64, target interface{}) error {
	_, err := ds.db.NewDelete().Model(target).Where("id = ?", id).Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DataStore) Update(id int64, target interface{}) error {
	_, err := ds.db.NewUpdate().Model(target).Where("id = ?", id).Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}
