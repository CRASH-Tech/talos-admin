package database

import (
	"time"

	v1 "github.com/CRASH-Tech/talos-admin/internal/talos-admin/api/v1"
)

func (db *DB) GetClusters() ([]v1.Cluster, error) {
	var clusters []v1.Cluster

	err := db.db.NewSelect().
		Model(&clusters).
		OrderExpr("name ASC").
		Limit(10).
		Scan(db.ctx)
	if err != nil {
		return clusters, err
	}

	return clusters, nil
}

func (db *DB) GetCluster(id int64) (v1.Cluster, error) {
	var cluster v1.Cluster

	err := db.db.NewSelect().
		Model(&cluster).
		Where("id = ?", id).
		Scan(db.ctx)
	if err != nil {
		return cluster, err
	}

	return cluster, nil
}

func (db *DB) AddCluster(c v1.Cluster) error {
	c.CreatedOn = time.Now()

	_, err := db.db.NewInsert().Model(&c).Exec(db.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteCluster(c v1.Cluster) error {
	_, err := db.db.NewDelete().Model(&c).WherePK().Exec(db.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateCluster(c v1.Cluster) error {
	_, err := db.db.NewUpdate().Model(&c).WherePK().Exec(db.ctx)
	if err != nil {
		return err
	}

	return nil
}
