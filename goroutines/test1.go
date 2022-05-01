package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main() started")

	c := make(chan string)
	go func() {
		fmt.Println(<-c)
	}()

	go func() {
		c <- "John"
		close(c)
	}()

	time.Sleep(time.Second)
	fmt.Println("main() stopped")
}

/*
Output:
	main() started
	John
	main() stopped
*/