package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//input: 1986-04-16T05:20:00+06:00
//output: Wed Apr 16 05:20:00 +0600 1986

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	var (
		t time.Time
		timeStr string
		err error
	)

	fmt.Scan(&timeStr)

	t, err = time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return err
	}

	fmt.Println(t.Format(time.UnixDate))
	return nil
}
