package sort

// 快速排序(自上而下，先分区，再处理子问题，即逐步确定各个pivot节点数据的位置)
/*
快排的思想是这样的：如果要排序数组中下标从 p 到 r 之间的一组数据，我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）。
我们遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放到中间。
经过这一步骤之后，数组 p 到 r 之间的数据就被分成了三个部分，
前面 p 到 q-1 之间都是小于 pivot 的，中间是 pivot，后面的 q+1 到 r 之间是大于 pivot 的。
*/

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	separatorSort(0, len(arr) - 1, arr)
}

func separatorSort(start, end int, arr []int) {
	if start >= end {
		return
	}

	i := partition(start, end, arr)
	separatorSort(start, i - 1, arr)
	separatorSort(i + 1, end, arr)
}

func partition(start, end int, arr []int) int {
	pivot := arr[end]

	i := start
	for j := start ; j < end ; j ++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i ++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}