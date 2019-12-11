package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{6,1,4,2,78,23,57,7}
	BubbleSort(arr)
	t.Log(arr)
}
