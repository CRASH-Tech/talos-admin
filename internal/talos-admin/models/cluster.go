package models

import (
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

type Cluster struct {
	bun.BaseModel `bun:"table:clusters"`
	ID            int64     `json:"id" bun:",pk,autoincrement"`
	Name          string    `json:"name"`
	CreatedOn     time.Time `json:"created_on"`
}

var (
	ErrNotFound = fmt.Errorf("resource could not be found")
)

func (ds *DataStore) GetClusters() ([]Cluster, error) {
	var clusters []Cluster

	err := ds.db.NewSelect().
		Model(&clusters).
		OrderExpr("name ASC").
		Limit(10).
		Scan(ds.ctx)
	if err != nil {
		return clusters, err
	}

	return clusters, nil
}

func (ds *DataStore) GetCluster(id int64) (Cluster, error) {
	var cluster Cluster

	err := ds.db.NewSelect().
		Model(&cluster).
		Where("id = ?", id).
		Scan(ds.ctx)
	if err != nil {
		return cluster, err
	}

	return cluster, nil
}

func (ds *DataStore) AddCluster(c Cluster) error {
	c.CreatedOn = time.Now()

	_, err := ds.db.NewInsert().Model(&c).Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DataStore) DeleteCluster(c Cluster) error {
	_, err := ds.db.NewDelete().Model(&c).WherePK().Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DataStore) UpdateCluster(c Cluster) error {
	_, err := ds.db.NewUpdate().Model(&c).WherePK().Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}
