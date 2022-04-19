package main

/*
	Given an array of integers nums and an integer target,
	return indices of the two numbers such that they add up to target.

	You may assume that each input would have exactly one solution,
	and you may not use the same element twice.
*/

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int) // key - index from nums, value - value from nums

	for i, elem := range nums {
		if index, ok := mp[target - elem]; ok {
			return []int{index, i}
		}
		mp[elem] = i
	}
	return []int{-1, -1}
}

func main() {
	arr := []int{4, 7, 11, 2, 15}
	fmt.Println(twoSum(arr, 9)) // [1, 3]
	// arr := []int{4, 7, 11, 2, 15, 5, 5}
	fmt.Println(twoSum(arr, 10)) // [5, 6]
}