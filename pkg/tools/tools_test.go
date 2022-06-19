package tools

import (
	"fmt"
	"log"
	"testing"
)

func TestMain(m *testing.M) {

	log.Println("=== Start run ping ===")
	result := (t_ping("1.1.1.1"))
	go Real_ping("8.8.8.8")
	Real_ping("1.1.1.1")
	fmt.Println(result.AvgRtt)
	log.Println("==================")
}

// ICMP PING 輸出值範例

// PING www.google.com (142.251.42.228):
// 24 bytes from 142.251.42.228: icmp_seq=0 time=35.858ms
// 24 bytes from 142.251.42.228: icmp_seq=1 time=32.654ms
// 24 bytes from 142.251.42.228: icmp_seq=2 time=32.735ms
// 24 bytes from 142.251.42.228: icmp_seq=3 time=32.819ms
// 24 bytes from 142.251.42.228: icmp_seq=4 time=32.701ms

// --- www.google.com ping statistics ---
// 5 packets transmitted, 5 packets received, 0% packet loss
// round-trip min/avg/max/stddev = 32.654ms/33.3534ms/35.858ms/1.253459ms
