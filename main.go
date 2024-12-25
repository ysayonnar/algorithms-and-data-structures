package main

import (
	ll "algos/dataStructures"
)

func main() {
	list := ll.LinkedList{}
	list.Append(12)
	list.Append(13)
	list.Append(42)
	list.InsertFront(666)
	list.InsertAfterValue(666, -12)
	list.InsertAfterValue(1412, 60)
	list.DeleteFirst()
	list.DeleteLast()
	list.PrintList()
}
