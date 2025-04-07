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

	if err != nil {
		fmt.Printf("something went wrong reading file %s: %v", inputFilePath, err)
		return
	}
	linesChan := getLinesChannel(f)

	for line := range linesChan {
		fmt.Printf("read: %s\n", line)
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
			_, err := f.Read(buffer)

			if err == io.EOF {
				return
			}
			parts := strings.Split(string(buffer), "\n")
			currentLine += parts[0]
			if len(parts) > 1 {
				ch <- currentLine
				currentLine = parts[len(parts)-1]
			}

		}
	}()
	return ch
}
