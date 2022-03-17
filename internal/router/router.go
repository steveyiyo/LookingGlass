package router

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
)

func HTTPServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	r.ParseForm()

	if r.Method == "POST" {
		Action := "NULL"
		IP := "NULL"
		for key, values := range r.Form {
			if key == "Action" {
				Action = values[0]
			}
			if key == "IP" {
				IP = values[0]
			}
		}

		if net.ParseIP(IP) == nil {
			export := "IP Address: '" + IP + "' - Invalid \n"
			io.WriteString(w, export)
			IP = "NULL"
		} else {
			export := "IP Address: '" + IP + "' - Valid \n"
			fmt.Printf(export)
			io.WriteString(w, export)

			if Action == "ping" {
				io.WriteString(w, "Running ping...\n\n")
				io.WriteString(w, ping(IP))
			}
			if Action == "mtr" {
				io.WriteString(w, "Running mtr...\n\n")
				io.WriteString(w, mtr(IP))
			}
			if Action == "routev4" {
				io.WriteString(w, "Checking BGP Route...\n\n")
				io.WriteString(w, routev4(IP))
			}
			if Action == "routev6" {
				io.WriteString(w, "Checking BGP Route...\n\n")
				io.WriteString(w, routev6(IP))
			}
		}

	}
}

func ping(IP string) string {
	fmt.Print("Running ping...\n")

	c := "ping -O -c 5 " + IP
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func mtr(IP string) string {
	fmt.Print("Running mtr...\n")

	c := "mtr -G 2 -c 5 -erwbz " + IP
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func routev4(IP string) string {
	fmt.Print("Checking BGP Route...\n")

	c := "vtysh -c 'show ip bgp " + IP + "'"
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func routev6(IP string) string {
	fmt.Print("Checking BGP Route...\n")

	c := "vtysh -c 'show bgp " + IP + "'"
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func CreateConnection() {
	http.HandleFunc("/", HTTPServer)
	err := http.ListenAndServe(":17286", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
