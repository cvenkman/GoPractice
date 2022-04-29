package main

/*
Требуется прочитать данные, и рассчитать среднее количество оценок,
полученное студентами группы. Ответ на задачу требуется записать
на стандартный вывод в формате JSON в следующей форме:
{
    "Average": 14.1
}
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type University struct {
	// ID int `json:"id"`
	// Number string `json:"Number"`
	// Year int `json:"Year"`
	Students []Student `json:"Students"`
}

type Student struct {
	// LastName string `json:"LastName"`
	// FirstName string `json:"FirstName"`
	// MiddleNam string `json:"MiddleName"`
	// Birthday string `json:"Birthday"`
	// Address string `json:"Address"`
	// Phone string `json:"Phone"`
	Rating []int `json:"Rating"`
}

type Average struct {
	Average float64
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	file, err := os.Open("students.json")
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

	university := University{}
	err = json.Unmarshal([]byte(data), &university)
	if err != nil {
		return err
	}

	average, err := countAverage(university)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n\n", average)

	return nil
}

func countAverage(university University) ([]byte, error) {
	var studentCount, ratingCount float64

	for _, student := range university.Students {
		studentCount++
		for range student.Rating {
			ratingCount++
		}
		// ratingCount += float64(len(student.Rating))
	}

	average := Average {
		Average: ratingCount / studentCount,
	}

	data, err := json.MarshalIndent(average, "", "    ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
