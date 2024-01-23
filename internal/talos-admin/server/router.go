package server

import (
	"log"
	"net/http"
	"time"

	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func setRouter(cfg config.Сonfig) *gin.Engine {
	router := gin.Default()
	controller, err := controller.New(cfg)
	if err != nil {
		log.Panic(err)
	}
	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	// Serve static files to frontend if server is started in production environment
	// if cfg.Env == "prod" {
	// 	router.Use(static.Serve("/", static.LocalFile("./assets/build", true)))
	// }

	router.Use(static.Serve("/", static.LocalFile("/home/crash/projects/CRASH-Tech/talos-admin/web/test-admin/dist", true)))
	//router.Use(static.Serve("/web", static.LocalFile("./web/talos-admin", true)))
	//router.Static("/web", "./web/talos-admin")

	// front := router.Group("/web")
	// front.Use()
	// {
	// 	router.Static("/", "./web/talos-admin")
	// }

	v1 := router.Group("/api/v1")
	v1.Use(customHeaders)
	v1.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "X-Total-Count"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	v1.Use(authorization)
	{
		clusters := v1.Group("/clusters")
		{
			clusters.GET("", controller.ClustersRead)
			clusters.GET(":id", controller.ClusterRead)
			clusters.POST("", controller.ClusterCreate)
			clusters.OPTIONS(":id", controller.ClusterOptions)
			clusters.DELETE(":id", controller.ClusterDelete)
			// accounts.PATCH(":id", c.UpdateAccount)
		}
		//...
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
