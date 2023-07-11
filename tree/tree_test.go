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
		value: 1,
		left: &Node{
			value: 2,
			left: &Node{
				value: 4,
			},
			right: &Node{
				value: 5,
			},
		},
		right: &Node{
			value: 3,
			right: &Node{
				value: 6,
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
