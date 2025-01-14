package main

import (
	"algos/leetcode"
	"fmt"
)

func main() {
	a := []int{2, 3, 1}
	b := []int{3, 1, 2}
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("Result: ", leetcode.FindThePrefixCommonArray(a, b))
}
