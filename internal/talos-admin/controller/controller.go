package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Controller struct {
	ds models.DataStore
}

func New(cfg config.Сonfig) (Controller, error) {
	var c Controller

	ds, err := models.NewDatastore(cfg)
	if err != nil {
		return c, err
	}

	//TODO: MAKE IT
	// err = ds.Init()
	// if err != nil {
	// 	return c, err
	// }

	c.ds = ds

	return c, nil
}

func (c *Controller) GetClusters(g *gin.Context) {
	clusters, err := c.ds.GetClusters()
	if err != nil {
		g.Header("X-Total-Count", "0")
		g.JSON(http.StatusInternalServerError, "")

		return
	}

	g.Header("X-Total-Count", fmt.Sprintf("%d", len(clusters)))

	g.JSON(http.StatusOK, clusters)
}

func (c *Controller) GetCluster(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		log.Error(err)
		g.Header("X-Total-Count", "0")
		g.JSON(http.StatusBadRequest, "")

		return
	}

	cluster, err := c.ds.GetCluster(int64(id))
	if err != nil {
		log.Error(err)
		g.Header("X-Total-Count", "0")
		g.JSON(http.StatusNotFound, "")

		return
	}

	g.Header("X-Total-Count", "1")
	g.JSON(http.StatusOK, cluster)
}
