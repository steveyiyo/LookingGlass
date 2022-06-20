package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/LookingGlass/internal/client"
	"github.com/steveyiyo/LookingGlass/internal/router"
	"github.com/steveyiyo/LookingGlass/internal/web"
)

func pageNotAvailable(c *gin.Context) {
	type notFound struct {
		Status  bool
		Message string
	}
	var notFoundObj notFound
	notFoundObj = notFound{false, "Page not found"}
	c.JSON(404, notFoundObj)
}

func Server() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/mtr/", client.MTR)
		apiv1.POST("/ping/", client.Ping)
		apiv1.POST("/bgp_route/", client.Route)
		apiv1.POST("/TestPING/", web.TestPING)
	}
	apiv1admin := r.Group("/api/v1/admin")
	{
		apiv1admin.POST("/AddPoP/", web.AddPoP)
	}
	r.NoRoute(pageNotAvailable)

	r.Run("127.0.0.1:32280")
}

func main() {
	if len(os.Args) > 1 {
		service := os.Args[1]

		switch service {
		case "server":
			Server()
		case "router":
			router.CreateConnection()
		default:
			log.Println("Please specify a service or client to run.")
		}
	} else {
		log.Println("Please specify a service or client to run.")
	}
}
