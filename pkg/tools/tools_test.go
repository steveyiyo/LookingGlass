package tools

import (
	"fmt"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Printf("Start run traceroute\n")
	fmt.Println(t_traceroute("103.172.124.1"))

	log.Printf("Start run ICMP ping\n")
	fmt.Println(t_ping("172.16.0.62"))

	// seconds := 4658000
	// fmt.Print(time.Duration(seconds) * time.Nanosecond)
}
