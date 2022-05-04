package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // fl
	fmt.Println(longestCommonPrefix([]string{"fl", "flow", "flight"})) // fl
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "fl"})) // fl
	fmt.Println(longestCommonPrefix([]string{"flower"})) //
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"})) //
	fmt.Println(longestCommonPrefix([]string{"x", "x", "x"})) // x
	fmt.Println(longestCommonPrefix([]string{"x"})) // x
	fmt.Println(longestCommonPrefix([]string{"flower", "dd", "flight"})) //
}

func longestCommonPrefix(strs []string) string {
    s := make([]byte, 0, 0)

	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	var ch byte
	for i := 0; i < minLen; i++ {
		ch = strs[0][i]
		for _, str := range strs {
			if str[i] != ch {
				return string(s)
			}
		}
		s = append(s, ch)
	}
	return string(s)
}