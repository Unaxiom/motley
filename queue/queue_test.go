package queue

import "testing"

func TestGenericQueue(t *testing.T) {
	var q BoundedQ
	q.SetMaxLength(0)
	q.Push(1)
	q.Push(2)

	if q.Len() != 2 {
		t.Fatal("All elements haven't been inserted")
	}
}

func TestQueueNegativeMaxLength(t *testing.T) {
	var q BoundedQ
	a := q.SetMaxLength(-1)
	if a.Error() != "SetMaxLength only accepts length >= 0. Provided length is -1" {
		t.Fatal(a.Error())
	}
}

func TestMaxLenAtBegin(t *testing.T) {
	var q BoundedQ
	q.SetMaxLength(2)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	if q.Len() != 2 {
		t.Fatal("Set Max Length test has failed. Max length should have been 2, but value is ", q.Len())
	}
}

func TestMaxLenTwice(t *testing.T) {
	var q BoundedQ
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.SetMaxLength(3)
	q.Push(5)
	q.Push(6)
	q.Push(7)
	if q.Len() != 3 {
		t.Fatal("Set Max Length Twice has failed. Value expected was 3, but received ", q.Len())
	}
}

func TestPushPop(t *testing.T) {
	var q BoundedQ
	q.SetMaxLength(3)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	var a interface{}
	a = q.Pop()
	if a.(int) != 2 {
		t.Fatal("Popped element wasn't the correct element. Expected 2, received ", a)
	}
	a = q.Pop()
	if a.(int) != 3 {
		t.Fatal("Popped element wasn't the correct element. Expected 3, received ", a)
	}
	a = q.Pop()
	if a.(int) != 4 {
		t.Fatal("Popped element wasn't the correct element. Expected 4, received ", a)
	}
	a = q.Pop()
	if a != nil {
		t.Fatal("Popped element wasn't the correct element. Expected nil, received ", a)
	}
}
