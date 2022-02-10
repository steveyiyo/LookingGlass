package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func AddPoP(c *gin.Context) {
	// TODO: Add PoP
	Authorized_Key := c.PostForm("Authorized_Key")
	PoP := c.PostForm("PoP")
	Router_IP := c.PostForm("MGMT_IP")

	if CheckIPValid(Router_IP) {

		uuid, err := uuid.NewUUID()
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		// Defind Json Type
		type List struct {
			POP_Name string
			UUID     string
		}
		type Info struct {
			UUID     string
			MGMT_IP  string
			Add_Time int
			Status   bool
		}

		type PoP_Json struct {
			POP_List []List
			POP_Info []Info
		}

		// TODO: Save PoP Information
		if Authentication(Authorized_Key) {

			var j PoP_Json

			// Add PoP List
			j.POP_List = append(j.POP_List, List{PoP, uuid.String()})

			// Add PoP Infomotion
			j.POP_Info = append(j.POP_Info, Info{UUID: uuid.String(), MGMT_IP: Router_IP, Add_Time: int(time.Now().Unix()), Status: true})

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
	} else {
		c.String(400, "Invalid IP")
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

func Authentication(Authorized_Key string) bool {
	// Loading .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check Authorized_Key
	if Authorized_Key == os.Getenv("Authorized_Key") {
		return true
	} else {
		return false
	}
}

func CheckIPValid(IP string) bool {
	if net.ParseIP(IP) == nil {
		return false
	} else {
		return true
	}
}
