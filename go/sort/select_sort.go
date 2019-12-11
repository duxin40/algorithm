package sort

// 选择排序
/*
选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。
但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
*/

func SelectSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		//查找最小值
		minIndex := i
		for j := i; j < len(arr); j ++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}