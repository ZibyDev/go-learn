package main

import "fmt"

const factor = 1000

const (
	Bps  = 1
	Kbps = Bps * factor
	Mbps = Kbps * factor
	Gbps = Mbps * factor
)

func main() {
	fmt.Println(Bps)
	fmt.Println(Kbps)
	fmt.Println(Mbps)
	fmt.Println(Gbps)
}
