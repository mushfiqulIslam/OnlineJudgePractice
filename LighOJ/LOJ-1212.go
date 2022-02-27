package main

import (
	"fmt"
)

type Node struct {
	value int
	head  *Node
	tail  *Node
}

type Deque struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Deque) Len() int {
	if q == nil {
		return 0
	}
	return q.length
}

func (q *Deque) PushFront(val int) {
	node := Node{
		value: val,
		tail:  q.head,
	}
	if q.length == 0 {
		q.head = &node
		q.tail = &node
	} else {
		q.head.head = &node
		q.head = &node
	}
	q.length++
}

func (q *Deque) PushBack(val int) {
	node := Node{
		value: val,
		head:  q.tail,
	}
	if q.length == 0 {
		q.head = &node
		q.tail = &node
	} else {
		q.tail.tail = &node
		q.tail = &node
	}
	q.length++
}

func (q *Deque) PopFront() int {
	if q.length <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	node := q.head
	if q.head.tail != nil {
		q.head.tail.head = nil
	}
	q.head = q.head.tail
	q.length--
	return node.value
}

func (q *Deque) PopBack() int {
	if q.length <= 0 {
		panic("deque: PopBack() called on empty queue")
	}
	node := q.tail
	if q.tail.head != nil {
		q.tail.head.tail = nil
	}
	q.tail = q.tail.head
	q.length--
	return node.value
}

func main() {
	var numOfCase, size, numOfCommand, value int
	var command string
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&size, &numOfCommand)
		q := Deque{}
		fmt.Printf("Case %d:\n", i)
		for j := numOfCommand; j > 0; j-- {
			fmt.Scan(&command)
			if q.Len() == size && (command == "pushRight" || command == "pushLeft") {
				fmt.Printf("The queue is full\n")
				fmt.Scan(&value)
			} else if command == "pushLeft" {
				fmt.Scan(&value)
				q.PushFront(value)
				fmt.Printf("Pushed in left: %d\n", value)
			} else if command == "pushRight" {
				fmt.Scan(&value)
				q.PushBack(value)
				fmt.Printf("Pushed in right: %d\n", value)
			} else if q.Len() == 0 {
				fmt.Printf("The queue is empty\n")
			} else if command == "popLeft" {
				value = q.PopFront()
				fmt.Printf("Popped from left: %d\n", value)
			} else {
				value = q.PopBack()
				fmt.Printf("Popped from right: %d\n", value)
			}
		}
	}
}
