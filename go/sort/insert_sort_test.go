package sort

import "testing"

func TestInsertSort(t *testing.T) {
	arr := []int{3,1,4,56,7,89,1}
	InsertSort(arr)
	t.Log(arr)
}
