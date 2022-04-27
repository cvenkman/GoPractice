package wf

import (
	"fmt"
	"log"
	"os"
)

func WriteFile() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	dataInBytes := []byte("first string\nsecond\nempty string later\n")
	file.Write(dataInBytes)
	
	data, err := os.ReadFile(file.Name())
	fmt.Println(string(data))
}
