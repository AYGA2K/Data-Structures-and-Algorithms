package linkedList

import (
	"fmt"
)

type Node struct {
	Val  any
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Display() {
	temp1 := l.Head
	fmt.Print("[")
	for temp1 != nil {
		fmt.Print(temp1.Val)
		if temp1.Next != nil {
			fmt.Print(",")
		}
		temp1 = temp1.Next
	}
	fmt.Print("] \n")
}

func (l *LinkedList) InsertAtHead(val any) {
	temp1 := &Node{val, nil}
	if l.Head == nil {
		l.Head = temp1
	} else {
		temp2 := l.Head
		l.Head = temp1
		temp1.Next = temp2
	}
	l.Length += 1
}

func (l *LinkedList) InsertAtTail(val any) {
	temp1 := &Node{val, nil}
	if l.Head == nil {
		l.Head = temp1
		l.Length += 1
		return

	} else {
		temp2 := l.Head
		for temp2.Next != nil {
			temp2 = temp2.Next
		}
		temp2.Next = temp1
		l.Length += 1

	}
}

func (l *LinkedList) InsertAtRandom(n int, val any) {
	if n == 0 {
		l.InsertAtHead(val)
		l.Length += 1
		return
	}
	if n == l.Length-1 {
		l.InsertAtTail(val)
		l.Length += 1
		return
	}
	temp1 := l.Head
	temp2 := &Node{val, nil}
	current_pos := 0
	for current_pos < n-1 {
		temp1 = temp1.Next
		current_pos++
	}
	temp2.Next = temp1.Next
	temp1.Next = temp2
	l.Length += 1
}

func (l *LinkedList) DeleteAtHead() {
	if l.Head == nil {
		return
	}
	temp := l.Head
	l.Head = temp.Next
	l.Length -= 1
}

func (l *LinkedList) DeleteAtTail() {
	if l.Head == nil {
		return
	}
	temp1 := l.Head
	var temp2 *Node
	for temp1.Next != nil {
		temp2 = temp1
		temp1 = temp1.Next
	}
	temp2.Next = nil
	l.Length -= 1

}

func (l *LinkedList) DeleteAtRandom(n int) {
	if n == 0 {
		l.DeleteAtHead()
		l.Length -= 1
		return
	}
	if n == l.Length-1 {
		l.DeleteAtTail()
		l.Length -= 1
		return
	}

	current_pos := 0
	temp1 := l.Head
	for current_pos < n-1 {
		current_pos++
		temp1 = temp1.Next
	}
	temp2 := temp1.Next
	temp1.Next = temp2.Next
	l.Length -= 1
}

func (l *LinkedList) ChangeValue(n int, val int) {
	temp := l.Head
	for i := 0; i < n; i++ {
		temp = temp.Next
	}
	temp.Val = val

}

func (l *LinkedList) Reverse() {
	for i := 0; i < l.Length-1; i++ {
		temp1 := l.Head
		for temp1.Next != nil {
			temp1 = temp1.Next
		}
		l.InsertAtRandom(i, temp1.Val)
		l.DeleteAtTail()

	}
}
