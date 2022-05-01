package main

import (
	"fmt"
)

/* Best Time to Buy and Sell Stock */

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{7, 2, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1})) // 0
	fmt.Println(maxProfit([]int{1, 3, 5, 6, 8})) // 7
	fmt.Println(maxProfit([]int{7, 8})) // 1
	fmt.Println(maxProfit([]int{7})) // 0
	fmt.Println(maxProfit([]int{2, 4, 1})) // 2
}

/* O(n) */
func maxProfit(prices []int) int {
	maxIndex := len(prices) - 1
	if maxIndex == 0 {
		return 0
	}

	// if isSorted(prices) == true {
	// 	if prices[maxIndex] - prices[0] > 0 {
	// 		return prices[maxIndex] - prices[0]
	// 	}
	// 	return 0
	// }

	max := 0
	left := 0
	right := 1
	for ; right <= maxIndex; right++ {
		if prices[right] < prices[left] {
			left = right
		} else if prices[right] - prices[left] > max {
			max = prices[right] - prices[left]
		}
	}
	return max
}

// func isSorted(prices []int) bool {
// 	g := prices[0]
	
// 	if prices[0] < prices[1] {
// 		/* если отсортированы по возрастанию */
// 		for _, price := range prices {
// 			if price < g {
// 				return false
// 			}
// 			g = price
// 		}
// 	} else {
// 		/* если отсортированы по убыванию */
// 		for _, price := range prices {
// 			if price > g {
// 				return false
// 			}
// 			g = price
// 		}
// 	}
// 	return true
// }
