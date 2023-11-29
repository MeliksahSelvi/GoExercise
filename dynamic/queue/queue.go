package main

import "fmt"

type MyItem int

type Queue struct {
	list []*MyItem
}

func NewQueue() *Queue {
	return &Queue{list: []*MyItem{}}
}

func (q *Queue) Enqueue(element MyItem) {
	q.list = append(q.list, &element)
}

func (q *Queue) Dequeue() MyItem {
	result := q.list[0]
	if result == nil {
		return -1
	}
	q.list = q.list[1:]
	return *result
}

func main() {
	q := NewQueue()

	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(*q)
}
