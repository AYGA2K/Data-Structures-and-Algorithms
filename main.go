package main

import (
	"fmt"
	"github.com/AYGA2K/DataStructuresAndAlgorithms/tree"
)

func main() {
	t := &tree.Node{Value: 5, Left: &tree.Node{Value: 10},
		Right: &tree.Node{Value: 20}}

	fmt.Println(tree.Bfs(t, 10))

}
