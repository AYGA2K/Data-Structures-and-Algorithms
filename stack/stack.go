package stack

type Node struct {
	value any
	next  *Node
}

type Stack struct {
	Head   *Node
	Tail   *Node
	Length int
}
