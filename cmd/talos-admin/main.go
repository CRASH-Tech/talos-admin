package main

import (
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/database"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/logger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// var (
// 	cfg config.СonfigImpl
// )

type App struct {
	cfg    config.Сonfig
	db     database.DB
	router *gin.Engine
	//log log.Logger
	//router XXX
}

// func init() {
// 	cfg = config.Get()

// 	// db := database.New(cfg)
// 	// db.Cluster().
// }

func NewApp(cfg config.Сonfig) *App {
	var app App

	logger.Init(cfg)
	app.cfg = cfg
	app.db = database.New(cfg)
	app.db.Init()

	return &app
}

func (a *App) Start() {
	log.Info("Start APP")
}

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Panic(err)
	}

	app := NewApp(cfg)
	app.Start()
	//fillDB(app)

	//GET CLUSTERS
	// clusters, err := app.db.GetClusters()
	// if err != nil {
	// 	log.Panic(err)
	// }
	// log.Debug(clusters)

	//GET CLUSTER
	// cluster1, err := app.db.GetCluster(1)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// log.Debug(cluster1)

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
	// c := v1.Cluster{
	// 	ID:   5,
	// 	Name: "lolll",
	// }
	// err = app.db.UpdateCluster(c)
	// if err != nil {
	// 	log.Panic(err)
	// }

	//server.Start(app.cfg)
}

func fillDB(a App) {
	// for i := 0; i < 10; i++ {
	// 	c := v1.Cluster{
	// 		Name: fmt.Sprintf("k-test-%d", i),
	// 	}
	// 	a.db.AddCluster(c)
	// }
}
