package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type myStruct struct {
		Name   string
		/* неэкспортируемые поля (имена которых начинаются со строчной буквы)
		не участвуют в кодировании / декодировании */
		age    int
		Status bool
		Values []int
		TestField []byte
	}
	
	s := myStruct{
		Name:   "John Connor",
		age:    35,
		Status: true,
		Values: []int{15, 11, 37},
		TestField: []byte("aa"),
	}
	
	/* Marshal */
	data, err := json.MarshalIndent(s, "", "\t")
	// data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n\n", data)

	/* Check JSON vlaid */
	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}

	/* Unarshal */
	s.Name = "dcs" // не изменит, т.к. Unmarshal возмет значения из data
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", s.Name)
}