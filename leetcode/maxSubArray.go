package main

/*
	Given an integer array nums, find the contiguous subarray
	(containing at least one number)
	which has the largest sum and return its sum.

	A subarray is a contiguous part of an array.
*/

import "fmt"

/* because math.Max works with float64 */
func findMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

/*
 * Kadane's alghorithm
 * Time - O(n)
 * Space - O(1)
*/
func maxSubArray(nums []int) int {
    sum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum = findMax(nums[i], nums[i] + sum)
		maxSum = findMax(maxSum, sum)
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
