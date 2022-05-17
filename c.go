package main

import (
	"fmt"
	"strconv"
	"log"
)

func main() {
	var N, n int
	var str string
	fmt.Scanln(&N)
	var err error

	arrr := make([][]int, N, N)


	// var m map[int]int
// занести в мапу где ключ - индекс, значение - число
..отсортир

	for i := 0; i < N; i++ {
		fmt.Scanln(&n)
		fmt.Scanln(&str)
		arr := make([]int, n, n)
		k := 0
		// fmt.Println(str, arr, n, len(str))
		for c := 0; c < len(str); c++ {
			if str[c] == ' ' {
				continue
			}
			fmt.Println(str[c])
			arr[k], err = strconv.Atoi(string(str[c]))
			if err != nil {
				log.Fatal(err)
			}
			k++
		}
		fmt.Println("arr2=", arr)

		// ar := make([]int, n, n)
		// max := findMax(arr)
		// for l := 0; l < n; l++ {
		// 	if arr[l] == max || arr[l] == max - 1 {
		// 		ar[l] = n
		// 	}
		// }
		// arrr[i] = ar
	}
	fmt.Println("-----------")

	for _, c := range arrr {
		fmt.Println(c)
	}
}

// func findMax(arr []int) int {
// 	max := arr[0]

// 	for _, num := range arr {
// 		if num > max {
// 			max = num
// 		}
// 	}
// 	return max
// }
