package sort

// 插入排序
/*
首先，我们将数组中的数据分为两个区间，已排序区间和未排序区间。对应两个区间的指针分别向左向右。
初始已排序区间只有一个元素，就是数组的第一个元素。
插入算法的核心思想是取未排序区间中的元素(向右)，在已排序区间中找到合适的插入位置将其插入(向左)，并保证已排序区间数据一直有序。
重复这个过程，直到未排序区间中元素为空，算法结束。
*/

func InsertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j + 1], arr[j] = arr[j], arr[j + 1]
			} else {
				break
			}
		}
		arr[j + 1] = value
	}
}