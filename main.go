package main

import (
	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/LookingGlass/internal/admin"
	"github.com/steveyiyo/LookingGlass/internal/client"
)

func pageNotAvailable(c *gin.Context) {
	type notFound struct {
		status  bool
		message string
	}
	var notFoundObj notFound
	notFoundObj = notFound{false, "Page not found"}
	c.JSON(404, notFoundObj)
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/mtr/", client.MTR)
		apiv1.POST("/ping/", client.Ping)
		apiv1.POST("/bgp_route/", client.Route)
	}
	apiv1admin := r.Group("/api/v1/admin")
	{
		apiv1admin.POST("/api/v1/admin/AddPoP/", admin.AddPoP)
	}
	r.NoRoute(pageNotAvailable)

	r.Run("127.0.0.1:32280")

}
