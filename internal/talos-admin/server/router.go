package server

import (
	"net/http"

	v1 "github.com/CRASH-Tech/talos-admin/internal/talos-admin/api/v1"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func setRouter(cfg config.СonfigImpl) *gin.Engine {
	router := gin.Default()

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

	rV1 := router.Group("/api/v1")
	rV1.Use(customHeaders)
	rV1.Use(authorization)
	{
		clusters := rV1.Group("/clusters")
		{
			clusters.GET(":id", v1.ClusterGet)
			clusters.GET("", v1.ClusterGetAll)
			// accounts.POST("", c.AddAccount)
			// accounts.DELETE(":id", c.DeleteAccount)
			// accounts.PATCH(":id", c.UpdateAccount)
			// accounts.POST(":id/images", c.UploadAccountImage)
		}
		//...
	}
	//////

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
