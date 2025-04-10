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
