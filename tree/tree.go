package tree

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

//---------------------------------------------------------//

func walk_preorder(node *Node, path []any) []any {
	if node != nil {
		path = append(path, node.Value)
		path = walk_preorder(node.Left, path)
		path = walk_preorder(node.Right, path)
	}
	return path
}

func PreOrderSearch(head *Node) []any {
	path := []any{}
	path = walk_preorder(head, path)
	return path
}

//---------------------------------------------------------//

func walk_inorder(node *Node, path []any) []any {
	if node != nil {
		path = walk_inorder(node.Left, path)
		path = append(path, node.Value)
		path = walk_inorder(node.Right, path)
	}
	return path
}

func InOrderSearch(head *Node) []any {
	path := []any{}
	path = walk_inorder(head, path)
	return path
}

//---------------------------------------------------------//

func walk_postorder(node *Node, path []any) []any {
	if node != nil {
		path = walk_postorder(node.Left, path)
		path = append(path, node.Value)
		path = walk_preorder(node.Right, path)
	}
	return path
}

func PostOrderSearch(head *Node) []any {
	path := []any{}
	path = walk_postorder(head, path)
	return path
}

// ---------------------------------------------------------//

func Bfs(head *Node, needle any) bool {
	queue := []*Node{head}

	for len(queue) > 0 {
		//Dequeue
		curr := queue[0]
		queue = queue[1:]
		if curr.Value == needle {
			return true
		}
		//Enqueue
		queue = append(queue, curr.Left)
		queue = append(queue, curr.Right)
	}

	return false
}

// ---------------------------------------------------------//

func Compare(t1 *Node, t2 *Node) bool {
	// we are going to use DFS
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Value != t2.Value {
		return false
	}
	return Compare(t1.Left, t2.Left) && Compare(t1.Right, t2.Right)
}

// ---------------------------------------------------------//

func search(curr *Node, needle int) bool {
	if curr == nil {
		return false
	}
	if curr.Value == needle {
		return true
	}
	if needle > curr.Value.(int) {
		search(curr.Right, needle)
	} else {

		search(curr.Right, needle)
	}
	return false
}

func Dfs(head *Node, needle int) bool {
	return search(head, needle)
}
