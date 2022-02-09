package admin

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tidwall/sjson"
)

const Init_Json = `{"POP_List":[],"POP_Info":[]}`

func AddPoP(c *gin.Context) {
	// Check Json File Available
	sjson.Set(Init_Json, "POP_List.POP_Name", "POP1")
	sjson.Set(Init_Json, "POP_List.uuid", "b8ddc37a-7fd9-4236-a54c-181cf3a1549c")

	// TODO: Add PoP
	Authodized_Key := c.PostForm("Authorized_Key")
	PoP := c.PostForm("PoP")
	Router_IP := c.PostForm("MGMT_IP")

	// TODO: Define PoP
	type PoP_Info struct {
		PoP string `json:"pop"`
		IP  string `json:"ip"`
	}

	// TODO: Save PoP Information
	if Authentication(Authodized_Key) {
		var PoP_Info_List []PoP_Info
		PoP_Info_List = append(PoP_Info_List, PoP_Info{PoP, Router_IP})
		c.String(200, "Success")
	} else {
		c.String(401, "Unauthorized")
	}
}

func CheckPoPAvailability(PoP string) bool {
	// Search Json
	return false
}

func Authentication(Authodized_Key string) bool {
	// Loading .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check Authodized_Key
	if Authodized_Key == os.Getenv("Authodized_Key") {
		return true
	} else {
		return false
	}
}
