package queue

import (
	"errors"
	"fmt"
	"sync"
)

// FixedLenQueue is a struct that acts as a queue (FIFO), with a limit imposed on the number of elements that it could hold. It can safely be used across goroutines.
type FixedLenQueue struct {
	maxLength int // Stores the maximum number of elements in the Queue
	list      []interface{}
	sync.RWMutex
}

// Push adds an element to the queue
func (q *FixedLenQueue) Push(element interface{}) {
	q.Lock()
	defer q.Unlock()
	q.push(element)
	// If maxLength is 0 (such as in default case, this would behave as a normal slice)
	if q.maxLength > 0 {
		q.prune()
	}
}

// push is the internal function that appends the element to the list without enabling any locks. Locks need to be handled at the outer level.
func (q *FixedLenQueue) push(element interface{}) {
	q.list = append(q.list, element)
}

// Pop removes and returns the first element from the queue
func (q *FixedLenQueue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()
	return q.pop()
}

// pop is the internal function that removes the first element of the list without usage of any locks. Locks need to be handled at the outer level.
func (q *FixedLenQueue) pop() interface{} {
	if q.len() == 0 {
		return nil
	}
	val := q.list[0]
	q.list = q.list[1:]
	return val
}

// Len returns the number of elements in the queue
func (q *FixedLenQueue) Len() int {
	q.RLock()
	defer q.RUnlock()
	return q.len()
}

// len returns the number of elements in the queue without usage of any locks
func (q *FixedLenQueue) len() int {
	return len(q.list)
}

// SetMaxLength would set the maximum size of the queue >=0. If 0, it would behave as a normal queue. If < 0, it would panic
func (q *FixedLenQueue) SetMaxLength(length int) error {
	// Call prune() here to trim the size of the queue after setting the maxLength attribute
	if length < 0 {
		err := fmt.Sprint("SetMaxLength only accepts length >= 0. Provided length is ", length)
		// panic(err)
		return errors.New(err)
	}
	q.Lock()
	defer q.Unlock()
	q.maxLength = length
	q.prune()
	return nil
}

// prune removes the first few elements such that the queue can be resized to the maxLength. Need to handle locks at a higher level.
func (q *FixedLenQueue) prune() {
	// q.Lock()
	// defer q.Unlock()
	for {
		if q.len() == 0 {
		}
		if q.len() > q.maxLength {
			q.list = q.list[1:]
		} else {
			return
		}
	}
}
