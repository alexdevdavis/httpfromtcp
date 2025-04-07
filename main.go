package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("something went wrong reading file: %v", err)
		return
	}
	byteStream := make([]byte, 8, 8)
	currentLine := ""
	for {
		_, err := f.Read(byteStream)
		if err == io.EOF {
			return
		}
		parts := strings.Split(string(byteStream), "\n")
		currentLine += parts[0]
		if len(parts) > 1 {
			fmt.Printf("read: %s\n", currentLine)
			currentLine = parts[1]
		}

	}
}
