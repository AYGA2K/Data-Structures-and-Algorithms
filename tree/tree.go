package tree

type Node struct {
	value any
	left  *Node
	right *Node
}

//---------------------------------------------------------//

func walk_preorder(node *Node, path []any) []any {
	if node != nil {
		path = append(path, node.value)
		path = walk_preorder(node.left, path)
		path = walk_preorder(node.right, path)
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
		path = walk_inorder(node.left, path)
		path = append(path, node.value)
		path = walk_inorder(node.right, path)
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
		path = walk_postorder(node.left, path)
		path = append(path, node.value)
		path = walk_preorder(node.right, path)
	}
	return path
}

func PostOrderSearch(head *Node) []any {
	path := []any{}
	path = walk_postorder(head, path)
	return path
}

//---------------------------------------------------------//
