package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
	len := len(nums) - 1

	for i := 0; i < len; i++ {
		if nums[i] == nums[i + 1] {
			return true
		}
	}
    return false
}

func main() {
	nums := []int{}
	fmt.Println(containsDuplicate(nums))
}