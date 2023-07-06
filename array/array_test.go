package array

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 2, 5}
	got := LinearSearch(arr, 2)
	fmt.Println("testing LinearSearch...")
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	} else {
		fmt.Println("\u2713")
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 5, 2, 7, 3, 4}
	BubbleSort(arr)
	got := arr
	fmt.Println("testing BubbleSort...")
	want := []int{1, 2, 3, 4, 5, 7}
	if !Compare(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	} else {
		fmt.Println("\u2713")
	}

}

func TestBinarySearch(t *testing.T) {
	fmt.Println("testing BinarySearch...")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := BinarySearch(arr, 8)
	want := true
	if got != want {
		t.Errorf("got %v , want %v", got, want)
	} else {
		fmt.Println("\u2713")
	}

}

func TestTwoCrystalBalls(t *testing.T) {
	fmt.Println("testing TwoCrystalBalls...")
	arr := []bool{false, false, false, false, false, true, true}
	got := TwoCrystalBalls(arr)
	want := 5
	if got != want {
		t.Errorf("got %v , want %v", got, want)
	} else {
		fmt.Println("\u2713")
	}

}
