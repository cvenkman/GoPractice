package main

/*
	Given an array nums of integers,
	return how many of them contain an even number of digits.
	1 <= nums.length <= 500
	1 <= nums[i] <= 10^5
*/

import (
	"fmt"
)

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896})) // 2
	fmt.Println(findNumbers([]int{12, 2100000000008899, 22, 1111, 7896})) // 5
	fmt.Println(findNumbers([]int{12})) // 1
	fmt.Println(findNumbers([]int{1})) // 0
}

func findNumbers(nums []int) int {
    count := 0

    for _, elem := range nums {
        i := 0;
        for ; elem > 0; i++ {
            elem /= 10
        }
        if i % 2 == 0 {
            count++
        }
    }
    return count
}