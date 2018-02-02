package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	// Ran this a few times, noticed that loopback and googleDNS came out in different orders.
	// Maps not ordered. change to remember from python dicts.
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
