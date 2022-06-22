package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/steveyiyo/LookingGlass/internal/agent"
	"github.com/steveyiyo/LookingGlass/internal/web"
)

func main() {
	godotenv.Load()
	webServerPort := os.Getenv("WebServerPort")

	if webServerPort == "" {
		log.Fatal("WebServerPort is not set")
	}

	if len(os.Args) > 1 {
		service := os.Args[1]

		switch service {
		case "server":
			web.WebServer(webServerPort)
		case "agent":
			agent.Init()
		default:
			log.Println("Please specify a service or client to run.")
		}
	} else {
		log.Println("Please specify a service or client to run.")
	}
}
