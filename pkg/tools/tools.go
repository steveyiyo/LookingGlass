package tools

import (
	"fmt"
	"net"

	"github.com/go-ping/ping"
)

func Real_ping(target string) string {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		panic(err)
	}

	var resultString string

	// 只進行 5 次 ping
	pinger.Count = 5

	pinger.OnRecv = func(pkt *ping.Packet) {
		data := fmt.Sprintf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		// fmt.Printf(data)
		resultString += data
	}

	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		data := fmt.Sprintf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
		// fmt.Printf(data)
		resultString += data
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		data := fmt.Sprintf("\n--- %s ping statistics ---\n", stats.Addr)
		// fmt.Printf(data)
		resultString += data

		data1 := fmt.Sprintf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		// fmt.Printf(data1)
		resultString += data1

		data2 := fmt.Sprintf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		// fmt.Printf(data2)
		resultString += data2
	}

	data := fmt.Sprintf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	// fmt.Printf(data)
	resultString += data

	err = pinger.Run()
	if err != nil {
		panic(err)
	}

	return resultString
}

func CheckIPValid(IP string) bool {
	if net.ParseIP(IP) == nil {
		return false
	} else {
		return true
	}
}
