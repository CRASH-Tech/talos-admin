package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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

func (db *DB) GetClusters() ([]Cluster, error) {
	var clusters []Cluster

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

func (db *DB) GetCluster(id int64) (Cluster, error) {
	var cluster Cluster

	err := db.db.NewSelect().
		Model(&cluster).
		Where("id = ?", id).
		Scan(db.ctx)
	if err != nil {
		return cluster, err
	}

	return cluster, nil
}

func (c *Cluster) Add() error {
	c.CreatedOn = time.Now()

	_, err := db.db.NewInsert().Model(&c).Exec(db.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cluster) Delete() error {
	_, err := db.db.NewDelete().Model(&c).WherePK().Exec(db.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cluster) Update() error {
	_, err := db.db.NewUpdate().Model(&c).WherePK().Exec(db.ctx)
	if err != nil {
		return err
	}

	return nil
}

// func GetAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

func ClusterGetAll(c *gin.Context) {
	//	c.Header("X-Total-Count", fmt.Sprintf("%d", len(clusters)))
	//
	// c.JSON(http.StatusOK, clusters)
}

func ClusterGet(c *gin.Context) {
	//id := c.Param("id")

	// for _, cluster := range clusters {
	// 	if cluster.ID == id {
	// 		c.Header("X-Total-Count", "1")
	// 		c.JSON(http.StatusOK, cluster)
	// 	}
	// }

	c.Error(ErrNotFound)
}
