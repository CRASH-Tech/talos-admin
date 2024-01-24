package controller

import (
	"fmt"
	"time"

	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/models"
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

	err = ds.Init()
	if err != nil {
		return c, err
	}

	c.ds = ds

	return c, nil
}

func (c *Controller) FillDB() {
	for i := 1; i < 10; i++ {
		cluster := models.Cluster{
			ID:        int64(i),
			Name:      fmt.Sprintf("k-test-%d", i),
			CreatedOn: time.Now(),
			UpdatedOn: time.Now(),
		}

		_, err := c.ds.Create(&cluster)
		if err != nil {
			log.Panic(err)
		}

		for x := 1; x < 5; x++ {
			node := models.Node{
				//ID:         int64(x),
				Name:       fmt.Sprintf("k-test-%d-%d", i, x),
				ClusterID:  int64(i),
				TemplateID: int64(i),
				CreatedOn:  time.Now(),
				UpdatedOn:  time.Now(),
			}

			_, err := c.ds.Create(&node)
			if err != nil {
				log.Panic(err)
			}
		}
	}

	for y := 1; y < 5; y++ {
		template := models.Template{
			//ID:         int64(x),
			Name:      fmt.Sprintf("template-%d", y),
			Data:      "{{ lol1 }}",
			CreatedOn: time.Now(),
			UpdatedOn: time.Now(),
		}

		_, err := c.ds.Create(&template)
		if err != nil {
			log.Panic(err)
		}
	}

	for z := 1; z < 5; z++ {
		variable := models.Variable{
			//ID:         int64(x),
			Name:      fmt.Sprintf("var-%d", z),
			Key:       fmt.Sprintf("KEY-%d", z),
			Value:     fmt.Sprintf("DATA-%d", z),
			CreatedOn: time.Now(),
			UpdatedOn: time.Now(),
		}

		_, err := c.ds.Create(&variable)
		if err != nil {
			log.Panic(err)
		}
	}
}
