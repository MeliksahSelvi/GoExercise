package main

import "fmt"

type MyItem int

type Stack struct {
	cap   uint64
	depth uint64
	list  []MyItem
}

func NewStack(cap uint64) *Stack {
	return &Stack{
		cap:   cap,
		depth: 0,
		list:  make([]MyItem, cap),
	}
}

func (s *Stack) Push(element MyItem) {

	if s.depth >= s.cap {
		panic("out of cap")
	}
	s.list[s.depth] = element
	s.depth++
}

func (s *Stack) Pop() MyItem {
	if s.depth < 1 {
		return -1
	}
	result := s.list[s.depth-1]
	s.list[s.depth-1] = 0
	s.depth--
	return result
}

func main() {

	s := NewStack(5)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
	println(s.Pop())
	fmt.Println(s)
}
