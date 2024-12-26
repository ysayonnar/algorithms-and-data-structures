package dataStructures

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (list *LinkedList) PrintList() {
	fmt.Print("List:\t")
	current := list.Head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Print("\n")
}

func (list *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}
	list.Tail.Next = newNode
	list.Tail = newNode
}

func (list *LinkedList) InsertFront(value int) {
	newNode := &Node{Value: value, Next: list.Head}

	//проверка на случай если список изначально пустой
	if list.Head == nil {
		list.Tail = newNode
	}
	list.Head = newNode
}

func (list *LinkedList) InsertAfterValue(afterValue int, value int) {
	newNode := &Node{Value: value, Next: nil}
	current := list.Head

	for current != nil {
		if current.Value == afterValue {
			newNode.Next = current.Next
			current.Next = newNode
			return
		}
		current = current.Next
	}

	//обработка кейса, если не найдено значение. Можно делать что угодно, но я решил сделать добавления элемента в конец
	list.Append(value)
}

func (list *LinkedList) DeleteFirst() {
	if list.Head != nil {
		//просто переназначаем голову, в си нужно было бы еще почистить предыдущий
		list.Head = list.Head.Next
	}
}

func (list *LinkedList) DeleteLast() {
	//если список пустой
	if list.Head == nil {
		return
	}

	//если в списке один элемент
	if list.Head == list.Tail {
		list.Head = nil
		list.Tail = nil
		return
	}

	//ищем элемент, предшествующий хвосту
	current := list.Head
	for current.Next != list.Tail {
		current = current.Next
	}

	current.Next = nil
	list.Tail = current
}
