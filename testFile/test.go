package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	if _, err := file.WriteString("first\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := file.WriteString("second\n"); err != nil {
		log.Fatal(err)
	}
	dataInBytes := []byte("[]byte string\n")
	file.Write(dataInBytes)
	
	/*	перезаписывает file.txt
		file2, _ := os.Create("file.txt")
		file2.WriteString("first1")
	*/

	var file2 *os.File
	file2, _ = os.Open("file.txt")
	file.WriteString("last")
	
	/*	"io/ioutil" is deprecated
		data := []byte("[]byte string")
		if err := ioutil.WriteFile(file.Name(), data, 0644); err != nil {
			log.Fatal(err)
		}
	*/
	
	data, err := os.ReadFile(file.Name())
	fmt.Println(string(data))
	
	res, _ := ioutil.ReadAll(file2)
	fmt.Println("-------")
	fmt.Println(string(res))
	file.Close()
}
