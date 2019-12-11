package sort

// 归并排序(自下而上，先排序子切片，再逐步合并)
/*
归并排序的核心思想还是蛮简单的。
如果要排序一个数组，我们先把数组从中间分成前后两部分，
然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了。
*/

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeSort(0, len(arr) - 1, arr)
}

func mergeSort(start, end int,  arr []int) {
	if start >= end {
		return
	}

	mid := (start + end)/2
	mergeSort(start, mid, arr)
	mergeSort(mid + 1, end, arr)
	merge(start, mid, end, arr)
}

func merge(start, mid, end int, arr []int) {
	tmp := make([]int, end - start + 1)

	i := start
	j := mid + 1
	k := 0
	for ;; k++ {
		if i > mid || j > end {
			break
		}
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i ++
		} else {
			tmp[k] = arr[j]
			j ++
		}
	}
	for ; i < mid + 1; i++ {
		tmp[k] = arr[i]
		k++
	}

	for ; j < end + 1; j++ {
		tmp[k] = arr[j]
		k++
	}

	copy(arr[start:end + 1], tmp)
}