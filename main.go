package main

import (
	"algos/leetcode"
	"fmt"
)

func main() {
	pings := []int{1, 100, 3001, 3002}
	obj := leetcode.Constructor()
	output := []int{}

	for _, item := range pings {
		output = append(output, obj.Ping(item))
	}

	fmt.Println("Output: ", output)
}
