package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type m struct {
	ID int `json:"global_id"`
}

type testStruct struct {
	Tests []struct {
	   Id int `json:"global_id"`
	} `json:"Tests"`
}

type MyStruct struct {
    GlobalID int `json:"global_id"`
}

type globalID struct {
	// description string
	// t int
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	globalIDstruct := new(testStruct)
	// globalIDstruct :
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if !json.Valid([]byte(data)) {
		fmt.Println("invalid json!")
		return err
	}
	err = json.Unmarshal(data, globalIDstruct)
	if err != nil {
		return err
	}
	// globalIDstruct=globalIDstruct

	// fmt.Println(string(data))

	var sum int

	for i := 0; i < len(globalIDstruct.Tests); i++ {
		sum += globalIDstruct.Tests[i].Id
	}

	fmt.Printf("%d", sum)
	return nil
}