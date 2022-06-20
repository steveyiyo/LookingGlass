package router

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/LookingGlass/pkg/tools"
)

const webServerPort = "59276"

func agent(c *gin.Context) {
	action := c.PostForm("action")
	target := c.PostForm("target")

	if !tools.CheckIPValid(target) {
		c.String(400, "Invalid IP")
	}

	switch action {
	case "ping":
		c.String(200, tools.Real_ping(target))
	case "mtr":
		c.String(404, "Not supported")
	case "routev4":
		c.String(404, "Not supported")
	case "routev6":
		c.String(404, "Not supported")
	default:
		c.String(400, "Invalid action")
	}
}

func WebServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.POST("/agent", agent)
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
	notFoundObj = notFound{false, "Not supported."}
	c.JSON(404, notFoundObj)
}

// func routev4(IP string) string {
// 	c := "vtysh -c 'show ip bgp " + IP + "'"
// 	cmd := exec.Command("bash", "-c", c)
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Println(err)
// 		return "An error occurred"
// 	}
// 	return string(out)
// }

// func routev6(IP string) string {
// 	c := "vtysh -c 'show bgp " + IP + "'"
// 	cmd := exec.Command("bash", "-c", c)
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Println(err)
// 		return "An error occurred"
// 	}
// 	return string(out)
// }
