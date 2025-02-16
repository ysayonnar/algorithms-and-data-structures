package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func generateRnd() int {
	return rand.Intn(100)
}

func main() {
	slots := make([][]int, 3)
	slots[0] = make([]int, 7)
	slots[1] = make([]int, 7)
	slots[2] = make([]int, 7)

	for i := 0; i < 7; i++ {
		slots[0][i] = generateRnd()
		slots[1][i] = generateRnd()
		slots[2][i] = generateRnd()
	}

	for i := 1; i < 50; i++ {
		ClearScreen()
		fmt.Printf("     %2d  %2d  %2d\n", slots[0][0], slots[1][0], slots[2][0])
		fmt.Printf("     %2d  %2d  %2d\n", slots[0][1], slots[1][1], slots[2][1])
		fmt.Printf("     %2d  %2d  %2d\n", slots[0][2], slots[1][2], slots[2][2])
		fmt.Printf("===> \033[32m%2d\033[0m  \033[32m%2d\033[0m  \033[32m%2d\033[0m\n", slots[0][3], slots[1][3], slots[2][3])
		fmt.Printf("     %2d  %2d  %2d\n", slots[0][4], slots[1][4], slots[2][4])
		fmt.Printf("     %2d  %2d  %2d\n", slots[0][5], slots[1][5], slots[2][5])
		fmt.Printf("     %2d  %2d  %2d\n", slots[0][6], slots[1][6], slots[2][6])

		delay := time.Duration(math.Log(float64(i+1))*100) * time.Millisecond
		if delay > 750*time.Millisecond {
			delay = 750 * time.Millisecond
		}
		time.Sleep(delay)

		slots[0] = slots[0][1:]
		slots[1] = slots[1][1:]
		slots[2] = slots[2][1:]

		slots[0] = append(slots[0], generateRnd())
		slots[1] = append(slots[1], generateRnd())
		slots[2] = append(slots[2], generateRnd())
	}

	if slots[0][3] == slots[1][2] && slots[0][2] == slots[2][2] && slots[1][2] == slots[2][2] {
		fmt.Println("ХОРОШ ПОБЕДА НАХУЙ")
	} else {
		fmt.Println("ФХАХФХАХАХФХХФХАХФХФХХАХАХВХАХВХА ПРОЕБАЛ ДАУН")
	}
}
