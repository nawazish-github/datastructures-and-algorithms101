package queue

type Node struct {
	Data int
	Next *Node
}

//Queue is the representation of the Queue data structure
type Queue struct {
	Head, Tail *Node
	Size       int
}

//Enqueue allows enqueing an element at the tail of the queue
func (q *Queue) Enqueue(elem int) {
	node := newNode(elem)
	//empty queue
	if q.Tail == nil && q.Head == nil {
		q.Tail = &node
		q.Head = &node
	}
	//when queue is non-empty
	q.Tail.Next = &node
	q.Tail = &node

	q.Size++
}

//Dequeue allows removal of the element from the head of the Queue
func (q *Queue) Dequeue() int {
	//when queue is empty
	if q.Head == nil {
		panic("Empty Queue")
	}
	//when queue is non-empty
	temp := q.Head
	q.Head = q.Head.Next

	q.Size--
	return temp.Data
}

func newNode(elem int) Node {
	return Node{
		Data: elem,
		Next: nil,
	}
}
