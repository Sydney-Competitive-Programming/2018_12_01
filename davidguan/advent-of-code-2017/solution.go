// go run advent-of-code-2017/solution.go < ../advent_of_code.txt gives "The digits sum is 1034"
package main

import (
	"fmt"
)

var ascIIZero = "0"[0]

func main() {
	s, sum := "", int64(0)
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		cur, next := s[i], s[(i+1)%len(s)]
		if cur == next {
			sum += int64(cur - ascIIZero)
		}
	}
	fmt.Println("The digits sum is", sum)
}
