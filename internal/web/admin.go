package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/steveyiyo/LookingGlass/pkg/tools"
)

func TestPING(c *gin.Context) {
	targetIP := c.PostForm("target")

	if tools.CheckIPValid(targetIP) {
		result := tools.Real_ping(targetIP)
		c.String(200, result)
	} else {
		c.String(400, "Failed")
	}

}

func AddPoP(c *gin.Context) {
	// TODO: Add PoP
	Authorized_Key := c.PostForm("Authorized_Key")
	PoP := c.PostForm("PoP")
	Router_IP := c.PostForm("MGMT_IP")

	type Message struct {
		Message string
	}

	if tools.CheckIPValid(Router_IP) {

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

		type POP_JSON struct {
			POP_List []List
			POP_Info []Info
		}

		// TODO: Save PoP Information
		if Authentication(Authorized_Key) {
			jsonFile, err := os.Open("pop_info.json")

			// if we os.Open returns an error then handle it
			if err != nil {
				log.Println(err)
			}

			// defer the closing of our jsonFile so that we can parse it later on
			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			var j POP_JSON

			json.Unmarshal([]byte(byteValue), &j)

			// Add PoP List
			j.POP_List = append(j.POP_List, List{PoP, uuid.String()})

			// Add PoP Infomotion
			j.POP_Info = append(j.POP_Info, Info{UUID: uuid.String(), MGMT_IP: Router_IP, Add_Time: int(time.Now().Unix()), Status: true})

			// Save to Json File
			b, err := json.Marshal(j)

			if err != nil {
				log.Println("json err:", err)
			}
			if SaveJsonFile(string(b)) {
				return_data := Message{Message: "Success"}
				c.JSON(200, return_data)
			} else {
				return_data := Message{Message: "Fail"}
				c.JSON(400, return_data)
			}
		} else {
			return_data := Message{Message: "Unauthorized"}
			c.JSON(401, return_data)
		}
	} else {
		return_data := Message{Message: "Failed"}
		c.JSON(400, return_data)
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
