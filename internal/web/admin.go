package web

import (
	"fmt"
	"log"
	"os"

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

func AgentJoin(c *gin.Context) {
	joinKey := c.Query("key")
	joinID := c.Query("id")

	// Query SQL

	// Add Data to SQL (Include PoP, UUID, MGMT_IP (Get it from gin IP) and Port, Add_Time, Status)

	log.Println(joinKey, joinID)
}

func AddPoP(c *gin.Context) {
	// TODO: Add PoP
	authorizedKey := c.PostForm("Authorized_Key")
	popName := c.PostForm("PoP")

	if Authentication(authorizedKey) {
		joinKey := tools.RandomString(16)
		joinID := uuid.New().String()
		agentJoinURL := fmt.Sprintf("http://%s/api/v1/admin/AgentJoin?key=%s&id=%s", c.Request.Host, joinKey, joinID)

		// 寫入 SQL
		log.Println(popName)

		// 確認寫入成功，回傳 200

		c.String(200, agentJoinURL)
	} else {
		resultMessage := Message{"Unauthorized"}
		c.JSON(400, resultMessage)
	}

	// if tools.CheckIPValid(Router_IP) {

	// 	uuid, err := uuid.NewUUID()
	// 	if err != nil {
	// 		fmt.Printf("%v\n", err)
	// 	}

	// 	// Defind Json Type
	// 	type List struct {
	// 		POP_Name string
	// 		UUID     string
	// 	}

	// 	type Info struct {
	// 		UUID     string
	// 		MGMT_IP  string
	// 		Add_Time int
	// 		Status   bool
	// 	}

	// 	type POP_JSON struct {
	// 		POP_List []List
	// 		POP_Info []Info
	// 	}

	// 	// TODO: Save PoP Information
	// 	if Authentication(Authorized_Key) {
	// 		jsonFile, err := os.Open("pop_info.json")

	// 		// if we os.Open returns an error then handle it
	// 		if err != nil {
	// 			log.Println(err)
	// 		}

	// 		// defer the closing of our jsonFile so that we can parse it later on
	// 		defer jsonFile.Close()

	// 		byteValue, _ := ioutil.ReadAll(jsonFile)

	// 		var j POP_JSON

	// 		json.Unmarshal([]byte(byteValue), &j)

	// 		// Add PoP List
	// 		j.POP_List = append(j.POP_List, List{PoP, uuid.String()})

	// 		// Add PoP Infomotion
	// 		j.POP_Info = append(j.POP_Info, Info{UUID: uuid.String(), MGMT_IP: Router_IP, Add_Time: int(time.Now().Unix()), Status: true})

	// 		// Save to Json File
	// 		b, err := json.Marshal(j)

	// 		if err != nil {
	// 			log.Println("json err:", err)
	// 		}
	// 		if SaveJsonFile(string(b)) {
	// 			return_data := Message{Message: "Success"}
	// 			c.JSON(200, return_data)
	// 		} else {
	// 			return_data := Message{Message: "Fail"}
	// 			c.JSON(400, return_data)
	// 		}
	// 	} else {
	// 		return_data := Message{Message: "Unauthorized"}
	// 		c.JSON(401, return_data)
	// 	}
	// } else {
	// 	return_data := Message{Message: "Failed"}
	// 	c.JSON(400, return_data)
	// }
}

// func SaveJsonFile(content string) bool {
// 	f, _ := os.Create("pop_info.json")
// 	f.WriteString(content)
// 	return true
// }

func CheckPoPAvailability(PoP string) bool {
	// Search Json
	// Idea: Create a loop to find the PoP?
	return false
}

func Authentication(authorizedKey string) bool {
	// Loading .env
	godotenv.Load()

	// Check Authorized_Key
	if authorizedKey == os.Getenv("Authorized_Key") {
		return true
	} else {
		return false
	}
}
