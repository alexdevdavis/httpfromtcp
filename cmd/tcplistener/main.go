package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	fmt.Println("server starting on port 42069...")
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Printf("error starting listener: %v\n", err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Printf("error accepting connection: %v\n", err)

		}
		fmt.Println("connection accepted")

		lc := getLinesChannel(c)
		for l := range lc {
			fmt.Println(l)
		}
		fmt.Println("closing connection")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		defer f.Close()
		buffer := make([]byte, 8, 8)
		currentLine := ""
		for {
			n, err := f.Read(buffer)

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file: %v", err)
				return
			}
			parts := strings.Split(string(buffer[:n]), "\n")
			for i := 0; i < len(parts)-1; i++ {
				currentLine += parts[i]
				ch <- currentLine
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
		if currentLine != "" {
			ch <- currentLine
		}
	}()
	return ch
}
