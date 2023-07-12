package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPreOrderSearch(t *testing.T) {
	// Create a sample tree
	//      1
	//     / \
	//    2   3
	//   / \   \
	//  4   5   6
	root := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
			},
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 3,
			Right: &Node{
				Value: 6,
			},
		},
	}

	want := []any{1, 2, 4, 5, 3, 6}

	got := PreOrderSearch(root)
	fmt.Println("testing PreOrderSearch...")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("PreOrderSearch() = %v, want %v", got, want)
	} else {
		fmt.Println("\u2713")

	}
}

func TestBfs(t *testing.T) {
	// Create a sample tree
	//      1
	//     / \
	//    2   3
	//   / \   \
	//  4   5   6
	root := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
			},
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 3,
			Right: &Node{
				Value: 6,
			},
		},
	}

	want := true

	got := Bfs(root, 3)
	fmt.Println("testing Bfs...")

	if got != want {
		t.Errorf("Bfs() = %v, want %v", got, want)
	} else {
		fmt.Println("\u2713")

	}
}

func TestCompare(t *testing.T) {

	// Create a sample tree
	//      1
	//     / \
	//    2   3
	//   / \   \
	//  4   5   6
	root := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
			},
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 3,
			Right: &Node{
				Value: 6,
			},
		},
	}
	root1 := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
			},
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 3,
			Right: &Node{
				Value: 6,
			},
		},
	}

	want := true

	got := Compare(root, root1)
	fmt.Println("testing Compare...")

	if got != want {
		t.Errorf("Compare() = %v, want %v", got, want)
	} else {
		fmt.Println("\u2713")

	}

}
