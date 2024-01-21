package v1

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
