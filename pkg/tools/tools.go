package tools

import (
	"fmt"

	"github.com/aeden/traceroute"
	"github.com/go-ping/ping"
	"github.com/nekomeowww/utillib/print"
)

func t_traceroute(target string) string {
	options := traceroute.TracerouteOptions{}

	c := make(chan traceroute.TracerouteHop, 0)
	go func() {
		for {
			_, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
		}
	}()

	result, _ := traceroute.Traceroute(target, &options, c)

	return print.SprintJSON(result)
}

func t_ping(target string) string {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		panic(err)
	}
	pinger.Count = 5
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats

	return print.SprintJSON(stats)
}
