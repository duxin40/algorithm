package sort

import "testing"

func TestSelectSort(t *testing.T) {
	arr := []int{2,1,4,35,67,24,3,2,999}
	SelectSort(arr)
	t.Log(arr)
}
