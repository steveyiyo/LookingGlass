package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func WebServer(webServerPort string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/mtr/", MTR)
		apiv1.POST("/ping/", Ping)
		apiv1.POST("/bgp_route/", Route)
		apiv1.POST("/TestPING/", TestPING)

		apiv1admin := apiv1.Group("/admin")
		{
			apiv1admin.POST("/AddPoP/", AddPoP)
		}
	}
	r.NoRoute(pageNotAvailable)

	ListenAddress := fmt.Sprintf("0.0.0.0:%s", webServerPort)
	log.Println("Starting web server on", ListenAddress)
	r.Run(ListenAddress)
}

func pageNotAvailable(c *gin.Context) {
	type notFound struct {
		Status  bool
		Message string
	}
	var notFoundObj notFound
	notFoundObj = notFound{false, "Page not found"}
	c.JSON(404, notFoundObj)
}
