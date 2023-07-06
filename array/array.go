package array

import "math"

func LinearSearch(haystack []int, needle int) bool {
	for _, num := range haystack {
		if num == needle {
			return true
		}
	}
	return false
}

func Compare(haystack_1 []int, haystack_2 []int) bool {
	for i, num := range haystack_1 {
		if num != haystack_2[i] {
			return false
		}
	}
	return true
}

func BubbleSort(haystack []int) {
	// O (N^2)
	for j := len(haystack); j > 0; j-- {
		for i := 0; i < j-1; i++ {
			if haystack[i] > haystack[i+1] {
				temp := haystack[i]
				haystack[i] = haystack[i+1]
				haystack[i+1] = temp
			}
		}
	}
}

func BinarySearch(haystack []int, needle int) bool {
	// [low,high[
	high := len(haystack)
	low := 0
	for low < high {
		m := low + (high-low)/2
		v := haystack[m]
		if v == needle {
			return true
		} else if needle > v {
			low = m + 1
		} else {
			high = m
		}
	}
	return false
}

func TwoCrystalBalls(breaks []bool) int {
	// O(sqr(N))
	jump_amount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jump_amount
	for ; i < len(breaks); i += jump_amount {
		if breaks[i] {
			break
		}
	}
	i -= jump_amount
	for j := 0; j < jump_amount && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}

	return -1
}
