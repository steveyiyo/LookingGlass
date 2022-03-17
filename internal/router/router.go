package router

import (
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
				io.WriteString(w, return_not_support())
				// io.WriteString(w, routev4(IP))
			}
			if Action == "routev6" {
				io.WriteString(w, "Checking BGP Route...\n\n")
				io.WriteString(w, return_not_support())
				// io.WriteString(w, routev6(IP))
			}
		}
	}
}

func ping(IP string) string {
	c := "ping -O -c 5 " + IP
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "An error occurred"
	}
	return string(out)
}

func mtr(IP string) string {
	c := "mtr -G 2 -c 5 -erwbz " + IP
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "An error occurred"
	}
	return string(out)
}

func routev4(IP string) string {
	c := "vtysh -c 'show ip bgp " + IP + "'"
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "An error occurred"
	}
	return string(out)
}

func routev6(IP string) string {
	c := "vtysh -c 'show bgp " + IP + "'"
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "An error occurred"
	}
	return string(out)
}

func return_not_support() string {
	return "Not support"
}

func CreateConnection() {
	http.HandleFunc("/", HTTPServer)
	err := http.ListenAndServe(":17286", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
