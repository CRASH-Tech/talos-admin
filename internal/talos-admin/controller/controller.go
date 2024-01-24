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

	var clusters []models.Cluster
	count, err := c.ds.ReadAll(&models.Cluster{}, &clusters, limit, offset, sortOrder)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")

		return
	}

	ctx.Header("X-Total-Count", strconv.Itoa(count))
	ctx.JSON(http.StatusOK, clusters)
}

func (c *Controller) ClusterRead(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, "")

		return
	}

	var cluster models.Cluster
	err = c.ds.Read(id, &cluster)
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

	id, err := c.ds.Create(&cluster)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")

		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (c *Controller) ClusterUpdate(ctx *gin.Context) {
	var cluster models.Cluster

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, "")

		return
	}

	cluster.ID = id

	err = ctx.ShouldBindJSON(&cluster)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, "")

		return
	}

	err = c.ds.Update(cluster.ID, &cluster)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")

		return
	}

	ctx.JSON(http.StatusOK, cluster)
}

func (c *Controller) ClusterDelete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, "")

		return
	}

	err = c.ds.Delete(id, &models.Cluster{})
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusNotFound, "")

		return
	}

	ctx.JSON(http.StatusOK, "")
}

func (c *Controller) ClusterOptions(ctx *gin.Context) {
	ctx.Header("Allow", "PUT,PATCH,OPTIONS,GET,DELETE")
	ctx.JSON(http.StatusOK, "")
}
