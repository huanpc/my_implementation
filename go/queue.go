package main

type Queue interface {
	IsEmpty() bool
	Enqueue(e interface{}) bool
	Dequeue() interface{}
	Length() int
}
