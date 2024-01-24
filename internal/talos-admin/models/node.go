package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Node struct {
	bun.BaseModel `bun:"table:nodes"`
	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	ClusterID     int64     `json:"cluster_id"`
	TemplateID    int64     `json:"template_id"`
	Name          string    `json:"name"`
	MachineConfig string    `json:"machine_config"`
	CreatedOn     time.Time `json:"created_on"`
	UpdatedOn     time.Time `json:"updated_on"`
}
