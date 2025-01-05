package main

import (
	"algos/leetcode"
	"fmt"
)

func main() {
	fmt.Println(leetcode.EqualPairs([][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	}))
}
