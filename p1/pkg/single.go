package pkg

import (
	"errors"
	"fmt"
)

type SingleQueue struct {
	maxSize int
	Queue [10]int
	front int
	rear int
}

func NewSingleQueue(size int) *SingleQueue {
	return &SingleQueue{
		maxSize: size,
		front: -1,
		rear: -1,
	}
}

func (q *SingleQueue) AddQueue(val int) (err error) {
	if q.rear == q.maxSize - 1 {
		return errors.New("The queue is full")
	}

	q.rear++	
	q.Queue[q.rear] = val

	return 
}

func (q *SingleQueue) ShowQueue() {
	if q.rear < 0 {
		fmt.Println("The queue is empty")
		return 
	}
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("Queue[%d]=%d\n", i, q.Queue[i])
	}
}