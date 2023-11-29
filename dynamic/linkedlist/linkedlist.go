package main

import "fmt"

type MyItem int

// tek yönlü çift yönlü olmasını istersek Prev Node'yi de saklamamız gerekir.
type Node struct {
	Data MyItem
	Next *Node
}

type LinkedList struct {
	Root *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(data MyItem) {
	n := &Node{
		Data: data,
		Next: nil,
	}

	if l.Root == nil {
		l.Root = n
		return
	}

	current := l.Root
	for current.Next != nil {
		current = current.Next
	}

	current.Next = n
}

func (l *LinkedList) Delete(data MyItem) {
	current := l.Root
	previous := l.Root

	for current != nil {

		if current.Data == data {
			previous.Next = current.Next
		}

		previous = current
		current = current.Next
	}
}

func (l *LinkedList) Print() {
	current := l.Root

	for current != nil {
		fmt.Printf(" -> %d", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	l := NewLinkedList()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Print()
	l.Delete(2)
	l.Print()
}
