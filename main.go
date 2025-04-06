package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("something went wrong reading file: %v", err)
		return
	}
	byteStream := make([]byte, 8)
	for {
		_, err := f.Read(byteStream)
		if err == io.EOF {
			return
		}
		fmt.Printf("read: %s\n", string(byteStream))
	}
}
