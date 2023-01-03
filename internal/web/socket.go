package web

import (
	"fmt"
	"log"
	"net"

	"github.com/valyala/fastjson"
)

func SocketServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Println(err)
	}

	log.Println("Server: TCP Server is running on port 8888")

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("Server: Accepted a new TCP connection from %s.", conn.RemoteAddr().String())

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	defer conn.Close()

	bytes := make([]byte, 1024)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			log.Println(err)
			return
		}
		receiveDate := string(bytes[:n])
		log.Printf("Client-`%s`: %s", conn.RemoteAddr().String(), receiveDate)
		if fastjson.Validate(receiveDate) != nil {
			log.Printf("Client-`%s`: Invalid JSON", conn.RemoteAddr().String())
			conn.Close()
			return
		}

		action := fastjson.GetString(bytes[:n], "ACTION")
		log.Println(action)
		switch action {
		case "register":
			// Check Auth Token
			checkAuth := checkAuthToken(fastjson.GetString(bytes[:n], "AUTH"))
			if !checkAuth {
				returnMessage := "Auth Token is invalid"
				conn.Write([]byte(fmt.Sprintf("%s\n", returnMessage)))
				log.Println(returnMessage)
				conn.Close()
				return
			}
			log.Println("Auth Token is valid")

			// Get POP Name
			popName := fastjson.GetString(bytes[:n], "POP_NAME")
			log.Printf("Register POP: %s", popName)

			// Save to Cache
			// If match, then do mtr

			returnMessage := "register"
			conn.Write([]byte(fmt.Sprintf("%s\n", returnMessage)))
			log.Println(returnMessage)
		}
	}
}
