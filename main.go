package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open("messages.txt")
	defer f.Close()
	if err != nil {
		fmt.Printf("something went wrong reading file %s: %v", inputFilePath, err)
		return
	}
	buffer := make([]byte, 8, 8)
	currentLine := ""
	for {
		_, err := f.Read(buffer)
		if err == io.EOF {
			return
		}
		parts := strings.Split(string(buffer), "\n")
		currentLine += parts[0]
		if len(parts) > 1 {
			fmt.Printf("read: %s\n", currentLine)
			currentLine = parts[1]
		}
	}
}
