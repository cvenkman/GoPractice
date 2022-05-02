package main

/*
	Given an integer array nums, move all the even integers at
	the beginning of the array followed by all the odd integers.

	Return any array that satisfies this condition.
	1 <= nums.length <= 5000
	0 <= nums[i] <= 5000
*/

import (
	"fmt"
)

func main() {
	fmt.Println(sortArrayByParitySecond([]int{3, 1, 2, 4})) // [2,4,3,1] or [4,2,3,1], [2,4,1,3], and [4,2,1,3]
	fmt.Println(sortArrayByParitySecond([]int{0})) // [0]
	fmt.Println(sortArrayByParitySecond([]int{0, 1, 4, 3, 6, 7})) // [0 4 6 7 3 1]
	fmt.Println(sortArrayByParitySecond([]int{13, 3, 11, 1, 3})) // [3 1 11 3 13]
	fmt.Println(sortArrayByParitySecond([]int{2, 6})) // [2 6]
}

/*  swap
	Time - O(n)
	Space - O(1)
*/
func sortArrayByParitySecond(nums []int) []int {
	if isAllOddOrEven(nums) {
		return nums
	}
	var j, i uint16
	j = uint16(len(nums) - 1)

	for tmp := 0; i <= j; {
		switch (nums[i] % 2) + (nums[j] % 2) {
		case 1:
			if nums[i] % 2 == 1 {
				tmp = nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			} else {
				i++
				j--
			}
		case 2:
			j--
		case 0:
			i++
		}
	}

	return nums

	/*  более понятное решение по такому же алгоритму:
		for tmp := 0; i <= j; {
			if nums[i] % 2 > nums[j] % 2 { // 3, 4 => 1 > 0
				tmp = nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			} else if nums[i] % 2 < nums[j] % 2 { // 4, 3 => 0 < 1
				i++
				j--
			} else if nums[i] % 2 == 1 && nums[j] % 2 == 1 {
				j--
			}  else if nums[i] % 2 == 0 && nums[j] % 2 == 0 {
				i++
			}
		}
	*/
}


/*  new array
	Time - O(n)
	Space - O(n)
*/
func sortArrayByParityFirst(nums []int) []int {
	if isAllOddOrEven(nums) {
		return nums
	}

	nums2 := make([]int, len(nums))
	var maxIndex, even, odd, i uint16
	maxIndex = uint16(len(nums) - 1)
	even = maxIndex

    for ; i <= maxIndex; i++ {
		if nums[i] % 2 == 0 { // четный
			nums2[odd] = nums[i]
			odd++
		} else {
			nums2[even] = nums[i]
			even--
		}
	}
	return nums2
}

func isAllOddOrEven(nums []int) bool {
	if nums[0] % 2 == 0 {
		for _, elem := range nums {
			if elem % 2 != 0 {
				return false
			}
		}
	} else if nums[0] % 2 != 0 {
		for _, elem := range nums {
			if elem % 2 == 0 {
				return false
			}
		}
	}
	return true
}