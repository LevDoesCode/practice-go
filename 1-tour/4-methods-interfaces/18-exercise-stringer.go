package main

import (
	"fmt"
	"strings"
)

type IPAddr [6]byte

func (ip IPAddr) String() string {
	var lel string
	// This is a more generalized form
	for current := range ip {
		lel = lel + fmt.Sprint(ip[current], " ")
	}
	return strings.Trim(lel, " ")
	// this only works for 4 bytes
	//return fmt.Sprint(ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1, 0, 0},
		"googleDNS": {8, 8, 8, 8, 0, 0},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
