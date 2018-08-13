package queue

import (
	"testing"

	"github.com/nawazish-github/datastructures-and-algorithms101/queue"
)

func TestEnqueue(t *testing.T) {
	var q queue.Queue
	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size != 2 {
		t.Errorf("Enque failed")
	}
}

func TestDeque(t *testing.T) {
	var q queue.Queue
	q.Enqueue(1)
	q.Enqueue(2)

	val := q.Dequeue()

	if val != 1 {
		t.Errorf("Dequeue failed to fetch head")
	}
	if q.Size != 1 {
		t.Errorf("Dequeue failed")
	}
}

func TestDequeForUnderFlow(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Underflow did not panic")
		}
	}()

	var q queue.Queue
	q.Enqueue(1)
	q.Enqueue(2)

	//Dequeue more than what was enqueued to make
	//queue underflow
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
}
