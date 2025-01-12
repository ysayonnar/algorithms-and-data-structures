package main

import (
	"algos/dataStructures"
	"fmt"
)

func main() {
	tree := dataStructures.Tree{}
	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(17)
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(-1)
	tree.Insert(75)
	tree.Insert(80)
	tree.Insert(90)
	tree.Insert(95)
	fmt.Println("Search 5: ", tree.Search(5))
	// tree.Delete(5)
	fmt.Println("Search 5: ", tree.Search(5))
	tree.PrintTreeStructure()
}
