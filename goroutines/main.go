package main

import (
	"fmt"
	"time"
)

func foo() {
	channel := make(chan string, 5)
	
	fmt.Println(len(channel), cap(channel)) // 0 5
	channel <- "cat"
	channel <- "dog"
	fmt.Println(len(channel), cap(channel)) // 2 5

	var str string
	str = <- channel
	fmt.Println(len(channel), cap(channel), str) // 1 5
	fmt.Println(len(channel), cap(channel), <- channel) // 1 5

	close(channel)

	val, ok := <- channel
	if ok == true {
		fmt.Println("--", len(channel), cap(channel), val) // 0 5
	}
}

func main() {
	start := time.Now()
	
	go foo()

	go func() {
		channel := make(chan int, 2)
		channel <- 2 + 2
	
		go func (c chan int) {
			fmt.Println("- ", <- c)
		}(channel)
	
		close(channel)
	}()

	duration := time.Since(start)
	time.Sleep(time.Second)
	fmt.Println(duration)
}

func deadlock() {
    fmt.Println("main() started")

	/*
	Deadlock:
		c := make(chan string)
		c <- "John"
	т.к. main - горутина, которая эксклюзивно производит операции с каналом.
	*/

    go func() {
		c := make(chan string)
		c <- "John"
	}()

    fmt.Println("main() stopped")
}