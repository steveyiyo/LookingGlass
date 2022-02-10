package client

import (
	"net"

	"github.com/gin-gonic/gin"
)

func MTR(c *gin.Context) {
	PoP := c.PostForm("PoP")
	IP := c.PostForm("Dst_IP")
	Action := "mtr"

	var result string
	if CheckIPValid(IP) {
		result = AssignedtoPoP(Action, PoP, IP)
		c.String(200, result)
	} else {
		result = "Invalid IP"
		c.String(400, result)
	}
}

func Ping(c *gin.Context) {
	PoP := c.PostForm("PoP")
	IP := c.PostForm("Dst_IP")
	Action := "ping"

	var result string
	if CheckIPValid(IP) {
		result = AssignedtoPoP(Action, PoP, IP)
		c.String(200, result)
	} else {
		result = "Invalid IP"
		c.String(400, result)
	}
}

func Route(c *gin.Context) {
	PoP := c.PostForm("PoP")
	IP := c.PostForm("Dst_IP")
	Action := "route"

	var result string
	if CheckIPValid(IP) {
		result = AssignedtoPoP(Action, PoP, IP)
		c.String(200, result)
	} else {
		result = "Invalid IP"
		c.String(400, result)
	}
}

func AssignedtoPoP(action string, PoP string, IP string) string {
	// Find PoP in PoP_List
	// Send HTTP Request to PoP (with Action)
	// Return Result
	return "Result"
}

func CheckIPValid(IP string) bool {
	if net.ParseIP(IP) == nil {
		return false
	} else {
		return true
	}
}
