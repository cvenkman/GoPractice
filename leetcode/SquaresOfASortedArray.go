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
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // 2
	// fmt.Println(sortedSquares([]int{12, 2100000000008899, 22, 1111, 7896})) // 5
	// fmt.Println(sortedSquares([]int{12})) // 1
	// fmt.Println(sortedSquares([]int{1})) // 0
}

func sortedSquares(nums []int) []int {
	tmp := 0
	maxIndex := len(nums) - 1
	fmt.Println(nums)
	count := 0
	
    for i := 0; count <= maxIndex; {
		sq := nums[i] * nums[i]
		if sq > nums[i + 1] {
			u := 0 
			for ; nums[u] < sq; u++ {

			}
			tmp = nums[u]
			nums[u] = sq
			nums

			// tmp = nums[maxIndex]
			// nums[maxIndex] = sq
			// nums[i] = tmp
		} else if sq < nums[maxIndex]  {
			// y := maxIndex
			// for ; nums[y] > sq; y-- {

			// }
			// tmp = nums[y]
			// nums[y] = sq
			// nums[i] = tmp
		} else {
			i++
		}
		count++
		fmt.Println(nums)
	}
	return nums
}