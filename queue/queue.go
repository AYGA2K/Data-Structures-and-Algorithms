package queue

type Node struct {
	Value interface{}
	Next  *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *Queue) Enque(val interface{}) {
	node := Node{Value: val}
	q.Length++

	if q.Tail == nil {
		q.Tail = &node
		q.Head = &node
		return
	}

	q.Tail.Next = &node
	q.Tail = &node

}

func (q *Queue) Deque() interface{} {
	if q.Head == nil {
		return nil
	}

	q.Length--
	head := q.Head
	q.Head = head.Next

	return head.Value
}

func (q *Queue) Peek() interface{} {

	if q.Head != nil {
		return q.Head.Value
	}

	return nil
}
