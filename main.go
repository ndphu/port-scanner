package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	TestHost = "portquiz.net"
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Elapse time: %s", elapsed)
}
func main() {
	fmt.Printf("Starting port scanning using host [%s]\n", TestHost)
	defer timeTrack(time.Now())
	wg := sync.WaitGroup{}

	for i := 1; i < 65536; i++ {
		go func(port int) {
			wg.Add(1)
			conn, err := net.Dial("tcp", TestHost+":"+strconv.Itoa(port))
			if err == nil {
				fmt.Printf("%d\n", port)
				conn.Close()
			}

			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Printf("Done\n")

}
