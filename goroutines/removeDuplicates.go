package main

import (
	"fmt"
	"time"
)

func main() {
	input := make(chan string, 3)
	output := make(chan string, 3)
	
	
	go func() {
		input <- "cat"
		input <- "dog"
		input <- "go"
	}()
	
	time.Sleep(time.Second * 2)

	go removeDuplicates(input, output)
}

func removeDuplicates(inputStream, outputStream chan string) {
	for {
		data, open := <- inputStream
		if !open {
			close(outputStream)
			return
		}
		fmt.Println(data)
	}
}