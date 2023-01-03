package web

import (
	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/LookingGlass/pkg/tools"
)

var popList []string

func checkAuthToken(key string) bool {
	return key == "authorized"
}

func checkPopExist(popName string) bool {
	for _, pop := range popList {
		if pop == popName {
			return true
		}
	}
	return false
}

func addPOP(popName string) bool {
	if !checkPopExist(popName) {
		popList = append(popList, popName)
		return true
	}
	return false
}

func removePOP(popname string) bool {
	for i := 0; i < len(popList); i++ {
		if popList[i] == popname {
			popList = append(popList[:i], popList[i+1:]...)
			i-- // form the remove item index to start iterate next item
		}
	}
	return true
}

func getPOPList() []string {
	return popList
}

func addActionTask(action, pop, ip string) string {
	// Generate a UUID
	return "123"
}

func queryResult(taskID string) string {
	return "ACK"
}

func MTR(c *gin.Context) {
	PoP := c.PostForm("PoP")
	IP := c.PostForm("Dst_IP")
	Action := "mtr"

	var result string
	if tools.CheckIPValid(IP) {
		taskID := addActionTask(Action, PoP, IP)
		go queryResult(taskID)
		resultData := Message{result}
		c.JSON(200, resultData)
	} else {
		result = "Invalid IP"
		resultData := Message{result}
		c.JSON(400, resultData)
	}
}

func Ping(c *gin.Context) {
	PoP := c.PostForm("PoP")
	IP := c.PostForm("Dst_IP")
	Action := "ping"

	var result string
	if tools.CheckIPValid(IP) {
		taskID := addActionTask(Action, PoP, IP)
		go queryResult(taskID)
		resultData := Message{result}
		c.JSON(200, resultData)
	} else {
		result = "Invalid IP"
		resultData := Message{result}
		c.JSON(400, resultData)
	}
}

func Route(c *gin.Context) {
	PoP := c.PostForm("PoP")
	IP := c.PostForm("Dst_IP")
	Action := "route"

	var result string
	if tools.CheckIPValid(IP) {
		taskID := addActionTask(Action, PoP, IP)
		go queryResult(taskID)
		resultData := Message{result}
		c.JSON(200, resultData)
	} else {
		result = "Invalid IP"
		resultData := Message{result}
		c.JSON(400, resultData)
	}
}
