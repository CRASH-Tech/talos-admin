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
	err = ds.Init()
	if err != nil {
		return c, err
	}

	c.ds = ds

	return c, nil
}

func (c *Controller) ClustersRead(ctx *gin.Context) {
	x := ctx.Request.URL.Query()
	var limit int
	var offset int

	starStr := x.Get("_start")
	endStr := x.Get("_end")
	if starStr != "" && endStr != "" {
		start, err := strconv.Atoi(starStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "")

			return
		}
		end, err := strconv.Atoi(endStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "")

			return
		}

		limit = end - start
		offset = start
	} else {
		limit = -1
		offset = -1
	}

	sort := x.Get("_sort")
	order := x.Get("_order")
	var sortOrder string
	if sort != "" && order != "" {
		sortOrder = fmt.Sprintf("%s %s", sort, order)
	}

	clusters, count, err := c.ds.ClustersRead(limit, offset, sortOrder)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")

		return
	}

	ctx.Header("X-Total-Count", strconv.Itoa(count))
	ctx.JSON(http.StatusOK, clusters)
}

func (c *Controller) ClusterRead(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, "")

		return
	}

	cluster, err := c.ds.ClusterRead(int64(id))
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusNotFound, "")

		return
	}

	ctx.JSON(http.StatusOK, cluster)
}

func (c *Controller) ClusterCreate(ctx *gin.Context) {
	var cluster models.Cluster

	err := ctx.ShouldBindJSON(&cluster)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, "")

		return
	}

	id, err := c.ds.ClusterCreate(cluster)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")

		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("%d", id))
}

func (c *Controller) ClusterDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, "")

		return
	}

	err = c.ds.ClusterDelete(int64(id))
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusNotFound, "")

		return
	}

	ctx.JSON(http.StatusOK, "")
}

func (c *Controller) ClusterOptions(ctx *gin.Context) {
	log.Info("OPTIONS")
	ctx.Header("Allow", "GET,POST,OPTIONS,DELETE")

	ctx.JSON(http.StatusOK, "")
}
