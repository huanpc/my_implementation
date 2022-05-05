// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	q := NewCircularQueue(2)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Enqueue(3)
	fmt.Println(q)
	fmt.Println(q.Peek())
	fmt.Println(q)
}

// CircularQueue defines a circular queue
type CircularQueue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
	count    int
}

// NewCircularQueue creates a queue
func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// IsEmpty returns true if queue is empty
func (q *CircularQueue) IsEmpty() bool {
	return q.count == 0
}

// IsFull returns true if queue is full
func (q *CircularQueue) IsFull() bool {
	return q.count == q.capacity
}

// Enqueue puts a element in the tail of queue
func (q *CircularQueue) Enqueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.count++
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

// Dequeue fetches a element from queue
func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	q.count--
	v := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

// String prints the queue
func (q *CircularQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := fmt.Sprintf("[head:%d]", q.head)
	var i = q.head
	for {
		result += fmt.Sprintf("<-%+v", q.data[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += fmt.Sprintf("<-[tail:%d]", q.tail)
	return result
}

func (q *CircularQueue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[q.head]
	return v
}
