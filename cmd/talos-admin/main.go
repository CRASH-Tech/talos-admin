package main

import (
	log "github.com/sirupsen/logrus"

	v1 "github.com/CRASH-Tech/talos-admin/internal/talos-admin/api/v1"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/database"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/logger"
)

// var (
// 	cfg config.СonfigImpl
// )

type App struct {
	cfg config.СonfigImpl
	db  database.DB
	//log log.Logger
	//router XXX
}

// func init() {
// 	// 	cfg = config.Get()

// 	// db := database.New(cfg)
// 	// db.Cluster().
// }

func (a *App) Init() *App {
	cfg := config.Get()
	logger.Init(cfg)
	a.cfg = cfg
	a.db = database.New(cfg)

	return a
}

func main() {
	app := App{}
	app.Init()

	//GET CLUSTERS
	clusters, err := app.db.GetClusters()
	if err != nil {
		log.Panic(err)
	}
	log.Debug(clusters)

	//GET CLUSTER
	cluster1, err := app.db.GetCluster(1)
	if err != nil {
		log.Panic(err)
	}
	log.Debug(cluster1)

	//ADD CLUSTER
	// c := v1.Cluster{
	// 	Name: "k-test-v",
	// }
	// err = app.db.AddCluster(c)
	// if err != nil {
	// 	log.Panic(err)
	// }

	//DELETE CLUSTER
	// c := v1.Cluster{
	// 	ID: 4,
	// }
	// err = app.db.DeleteCluster(c)
	// if err != nil {
	// 	log.Panic(err)
	// }

	//UPDATE CLUSTER
	c := v1.Cluster{
		ID:   5,
		Name: "lolll",
	}
	err = app.db.UpdateCluster(c)
	if err != nil {
		log.Panic(err)
	}

	//server.Start(app.cfg)
}
