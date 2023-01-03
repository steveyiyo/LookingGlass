package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string
}

func WebServer(webServerPort string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/pop_list/", getPOP)
		apiv1.POST("/mtr/", MTR)
		apiv1.POST("/ping/", Ping)
		apiv1.POST("/bgp_route/", Route)
	}
	r.NoRoute(pageNotAvailable)

	ListenAddress := fmt.Sprintf("0.0.0.0:%s", webServerPort)
	log.Println("Starting web server on", ListenAddress)
	r.Run(ListenAddress)
}

func getPOP(c *gin.Context) {
	popData := getPOPList()
	type POPList struct {
		PoP []string
	}
	c.JSON(200, POPList{popData})
}

func pageNotAvailable(c *gin.Context) {
	type notFound struct {
		Status  bool
		Message string
	}
	notFoundObj := notFound{false, "Page not found"}
	c.JSON(404, notFoundObj)
}
