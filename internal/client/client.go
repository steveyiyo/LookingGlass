package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

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
	// Defind Json Type
	type POPListItem struct {
		POP_Name string
		UUID     string
	}

	type POPInfoItem struct {
		UUID     string
		MGMT_IP  string
		Add_Time int
		Status   bool
	}

	type POP_JSON struct {
		POPList []POPListItem `json:"POP_List"`
		POPInfo []POPInfoItem `json:"POP_Info"`
	}

	var jt POP_JSON

	// Open Json File
	jsonFile, err := os.Open("pop_info.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &jt)

	var MGMT_IP string
	// Find PoP in POP_List
	for _, item := range jt.POPList {
		// Find PoP in POP_Info
		for _, item2 := range jt.POPInfo {
			// match UUID and Name
			if item.POP_Name == PoP && item2.UUID == item.UUID {
				// Find PoP in PoP_Info
				MGMT_IP = item2.MGMT_IP
			}
		}
	}

	// Check if PoP is found
	if MGMT_IP != "" {
		// Send TCP Request to PoP (with Action)
		conn, err := net.Dial("tcp", MGMT_IP+":17286")
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()
		conn.Write([]byte(action + " " + IP))
		message, _ := ioutil.ReadAll(conn)
		return string(message)
	} else {
		return "PoP not found"
	}
}

func CheckIPValid(IP string) bool {
	if net.ParseIP(IP) == nil {
		return false
	} else {
		return true
	}
}
