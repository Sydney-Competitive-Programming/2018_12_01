// DFS to search all the possibilities, which can be quite long.......
// Still running....
package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func validate(src []int, n int) bool {
	expectedSum := n
	for i := 2; i < n; i++ {
		expectedSum *= i
	}
	guess := n*n + 10

	if len(src) >= (expectedSum*(n-n/2)+1) || len(src) >= guess {
		return true
	}
	sum := 0
	indexList := []int{}
	for i := n - 1; i < len(src); i++ {
		checker := map[int]bool{}
		for j := 0; j < n; j++ {
			checker[src[i-j]] = true
		}

		if len(checker) == n {
			duplicated := false
			for _, index := range indexList {
				if testEq(src[index-n+1:index+1], src[i-n+1:i+1]) {
					duplicated = true
				}
			}
			if duplicated {
				continue
			}
			indexList = append(indexList, i)
			sum++
		}
	}
	return sum == expectedSum
}

// The potential coming task that for a number cur lived in index level can work on.
// func genTasks(n, cur, index int) [][2]int {
// 	tasks := [][2]int{}
// 	for i := 1; i <= n; i++ {
// 		if i == cur {
// 			continue
// 		}
// 		tasks = append(tasks, [2]int{index + 1, i})
// 	}
// 	return tasks
// }

func genTasks(n, index, cur int, state []int) [][2]int {
	tasks := [][2]int{}
	M := map[int]int{}
	for i := index; i > (index - n); i-- {
		M[state[i]]++
	}
	lastHalf := map[int]bool{}
	for i := index; i > (index - n/2); i-- {
		lastHalf[i] = true
	}
	for i := 1; i <= n; i++ {
		if (M[i] > 1) || lastHalf[i] {
			continue
		}
		tasks = append(tasks, [2]int{index + 1, i})
	}
	return tasks
}

func dfs(state []int, n int) int {
	q, shortest := genTasks(n, len(state)-1, state[n-1], state), math.MaxInt32

	for len(q) > 0 {
		task := q[0]
		q = q[1:]
		if len(state) > task[0] {
			state = append(state[:task[0]], task[1])
		} else if len(state) == task[0] {
			state = append(state, task[1])
		} else {
			fmt.Println(state, task)
			log.Fatal("should not happen, panic")
		}

		done := validate(state, n)

		if done {
			if shortest > len(state) {
				fmt.Println(state)
				shortest = len(state)
			}
		} else {
			q = append(genTasks(n, task[0], task[1], state), q...)
		}
	}
	return shortest
}

func main() {
	n := 0
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println(n)
		os.Exit(0)
	}
	initalState := make([]int, n)
	for i := 0; i < n; i++ {
		initalState[i] = i + 1
	}
	shortestLength := dfs(initalState, n)
	fmt.Println(shortestLength)
}
