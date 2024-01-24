package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Template struct {
	bun.BaseModel `bun:"table:templates"`
	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	Name          string    `json:"name"`
	Data          string    `json:"data"`
	CreatedOn     time.Time `json:"created_on"`
	UpdatedOn     time.Time `json:"updated_on"`
}
