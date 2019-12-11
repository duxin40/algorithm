package sort

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{2}
	MergeSort(arr)
	t.Log(arr)

	arr = []int{3,1}
	MergeSort(arr)
	t.Log(arr)

	arr = []int{1,8,5,23,47,2}
	MergeSort(arr)
	t.Log(arr)
}
