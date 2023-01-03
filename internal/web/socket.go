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

	registerStatus := false
	popName := ""
	defer func() {
		log.Println("Server: Connection closed.")
		if registerStatus {
			removePOP(popName)
			log.Printf("We've remove %s from POP list.", popName)
		}
		conn.Close()
	}()

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
		if action == "" {
			log.Printf("Client-`%s`: Invalid Action", conn.RemoteAddr().String())
			conn.Close()
			return
		}

		log.Println(action)
		switch action {
		case "register":
			// Check Auth Token
			checkAuth := checkAuthToken(fastjson.GetString(bytes[:n], "AUTHORIZED_KEY"))
			if !checkAuth {
				returnMessage := "Auth Token is invalid"
				conn.Write([]byte(fmt.Sprintf("%s\n", returnMessage)))
				log.Println(returnMessage)
				conn.Close()
				return
			}
			log.Println("Auth Token is valid")

			// Get POP Name
			popName = fastjson.GetString(bytes[:n], "POP_NAME")
			if popName == "" {
				returnMessage := "POP Name is invalid"
				conn.Write([]byte(fmt.Sprintf("%s\n", returnMessage)))
				log.Println(returnMessage)
				conn.Close()
				return
			}

			// Save to List
			if !addPOP(popName) {
				return
			}

			registerStatus = true
			returnMessage := fmt.Sprintf("Register POP: %s, ACK!", popName)
			conn.Write([]byte(fmt.Sprintf("%s\n", returnMessage)))
			log.Println(returnMessage)

		default:
			returnMessage := "Invalid Action"
			conn.Write([]byte(fmt.Sprintf("%s\n", returnMessage)))
			log.Println(returnMessage)
		}
	}
}
