package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/tidwall/sjson"
)

const Init_Json = `{"POP_List":[],"POP_Info":[]}`

func AddPoP(c *gin.Context) {
	// TODO: Add PoP
	Authodized_Key := c.PostForm("Authorized_Key")
	PoP := c.PostForm("PoP")
	Router_IP := c.PostForm("MGMT_IP")

	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Defind Json Type
	type List struct {
		POP_Name string
		uuid     string
	}
	type Info struct {
		uuid     string
		mgmt_IP  string
		Add_Time int
		status   bool
	}

	type PoP_Json struct {
		POP_List []List
		POP_Info []Info
	}

	// TODO: Save PoP Information
	if Authentication(Authodized_Key) {

		var j PoP_Json

		// Add PoP List
		j.POP_List = append(j.POP_List, List{PoP, uuid.String()})
		sjson.Set(Init_Json, "POP_List.POP_Name", PoP)
		sjson.Set(Init_Json, "POP_List.uuid", uuid)

		// Add PoP Infomotion
		j.POP_Info = append(j.POP_Info, Info{uuid: uuid.String(), mgmt_IP: Router_IP, Add_Time: int(time.Now().Unix()), status: true})

		// Save to Json File
		b, err := json.Marshal(j)

		if err != nil {
			fmt.Println("json err:", err)
		}
		if SaveJsonFile(string(b)) {
			c.String(200, "Success")
		} else {
			c.String(400, "Fail")
		}
	} else {
		c.String(401, "Unauthorized")
	}
}

func SaveJsonFile(content string) bool {
	f, _ := os.Create("pop_info.json")
	f.WriteString(content)
	return true
}

func CheckPoPAvailability(PoP string) bool {
	// Search Json
	// Idea: Create a loop to find the PoP?
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
