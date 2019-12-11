package sort

//冒泡排序
/*
冒泡排序只会操作相邻的两个数据。
每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。如果不满足就让它俩互换。
一次冒泡会让至少一个元素移动到它应该在的位置，
重复 n 次，就完成了 n 个数据的排序工作。
*/

func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i :=0; i < len(arr); i ++ {
		changed := false
		for j := 0; j < len(arr) - i - 1; j ++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}
