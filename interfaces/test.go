// package main

// import (

// )

// func main() {

// }

package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

	num := 1
	fmt.Print(num, " ")
	num += num

	for ; num <= n; num += num {
		fmt.Print(num, " ")
	}
}