package queue

type Node struct {
	value any
	next  *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *Queue) Enque(val any) {
	node := Node{value: val}
	q.Length++
	if q.Tail == nil {
		q.Tail = &node
		q.Head = &node
		return
	}
	q.Tail.next = &node
	q.Tail = &node

}
func (q *Queue) Deque() any {
	if q.Head == nil {
		return nil
	}
	q.Length--
	head := q.Head
	q.Head = head.next

	return head.value

}
func (q *Queue) Peek() any {
	if q.Head != nil {
		return q.Head.value
	}
	return nil
}
