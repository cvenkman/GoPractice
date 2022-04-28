package main

import (
	"fmt"
	"time"
)

func foo() {
	time.Sleep(time.Second)
	fmt.Println("foo")
}

/*
При отправлении значения в небуферизованный канал отправляющая горутина блокируется до тех пор,
пока другая горутина не выполнит получение из этого канала. После этого обе горутины продолжают работать.
Верно и обратное, если горутина получает значение, она блокируется, пока значение не будет получено.
*/

func main() {
	start := time.Now()
	
	// time.Sleep(time.Second)
	// fmt.Println("main")
	// go func() {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("goroutine")
	// } ()
	channel := make(chan string, 5)
	// num = <- channel
	go fmt.Println(len(channel), cap(channel)) // 0 5
	channel <- "cat"
	channel <- "dog"
	go fmt.Println(len(channel), cap(channel)) // 1 5
	var str string
	str = <- channel
	go fmt.Println(len(channel), cap(channel), str) // 0 5
	go fmt.Println(len(channel), cap(channel), <- channel) // 0 5
	close(channel)

	val, ok := <- channel
	if ok == true {
		fmt.Println("--", len(channel), cap(channel), val) // 0 5
	} 

	// go foo()
	duration := time.Since(start)
	time.Sleep(time.Second)
	fmt.Println(duration)
}