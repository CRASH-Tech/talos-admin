package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Variable struct {
	bun.BaseModel `bun:"table:variables"`
	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	ClusterID     int64     `json:"cluster_id"`
	Name          string    `json:"name"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
	CreatedOn     time.Time `json:"created_on"`
	UpdatedOn     time.Time `json:"updated_on"`
}
