package main

import (
	"fmt"
	"strings"
)

type Node struct {
	val  interface{}
	next *Node
}

type LinkedListQueue struct {
	head *Node // sentinal node
	tail *Node
	size int
}

func NewArrayQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head: &Node{},
	}
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.head.next == nil
}

func (q *LinkedListQueue) Enqueue(e interface{}) bool {
	tmp := &Node{val: e}
	if q.tail == nil {
		q.tail = tmp
		q.head.next = q.tail
	} else {
		q.tail.next = tmp
		q.tail = tmp
	}
	q.size++
	return true
}

func (q *LinkedListQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	e := q.head.next
	q.head.next = e.next
	if q.head.next == nil { // empty
		q.tail = nil
	}
	q.size--
	return e.val
}

func (q *LinkedListQueue) Len() int {
	return q.size
}

func (q *LinkedListQueue) String() string {
	ls := make([]string, 0, q.Len())
	for !q.IsEmpty() {
		e := q.Dequeue()
		ls = append(ls, fmt.Sprintf("%v", e))
	}
	return strings.Join(ls, "<-")
}

func main() {
	q := NewArrayQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	fmt.Println(q.String())
}
