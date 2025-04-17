package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Printf("error resolving udp address: %v\n", err)
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n> ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading input: %v ...try again\n", err)
		} else {
			_, err := conn.Write([]byte(s))
			if err != nil {
				fmt.Printf("error sending over udp: %v\n", err)
			}
		}
	}
}
