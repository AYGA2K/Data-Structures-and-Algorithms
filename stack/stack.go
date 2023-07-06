package stack

type Node struct {
	value any
	prev  *Node
}

type Stack struct {
	Head   *Node
	Length int
}

func (s *Stack) Posh(val any) {
	node := Node{value: val}
	s.Length++
	if s.Head == nil {
		s.Head = &node
		return
	}
	head := s.Head
	s.Head = &node
	s.Head.prev = head
}

func (s *Stack) Peek() any {
	if s.Head != nil {
		return s.Head.value
	}
	return nil
}
