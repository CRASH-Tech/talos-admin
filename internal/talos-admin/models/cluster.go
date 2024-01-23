package models

import (
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

type Cluster struct {
	bun.BaseModel `bun:"table:clusters"`
	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	Name          string    `json:"name"`
	CreatedOn     time.Time `json:"created_on"`
}

var (
	ErrNotFound = fmt.Errorf("resource could not be found")
)

func (ds *DataStore) ClustersRead(limit, offset int, sort string) ([]Cluster, int, error) {
	var clusters []Cluster

	q := ds.db.NewSelect().Model(&clusters)

	if limit >= 0 {
		q = q.Limit(limit)
	}

	if offset >= 0 {
		q = q.Offset(offset)
	}

	if sort != "" {
		q = q.OrderExpr(sort)
	}

	count, err := q.ScanAndCount(ds.ctx)
	if err != nil {
		return clusters, 0, err
	}

	return clusters, count, nil
}

func (ds *DataStore) ClusterRead(id int64) (Cluster, error) {
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

func (ds *DataStore) ClusterCreate(c Cluster) error {
	c.CreatedOn = time.Now()

	_, err := ds.db.NewInsert().Model(&c).Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DataStore) ClusterDelete(id int64) error {
	//TODO: CHECK IT
	//var cluster Cluster

	//cluster.ID = id
	_, err := ds.db.NewDelete().Model(Cluster{}).Where("id = ?", id).Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DataStore) ClusterUpdate(c Cluster) error {
	_, err := ds.db.NewUpdate().Model(&c).WherePK().Exec(ds.ctx)
	if err != nil {
		return err
	}

	return nil
}
