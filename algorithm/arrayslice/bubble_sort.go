package arrayslice

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/21 00:19
 * @Desc:	冒泡排序：
			外层循环控制需要进行多少轮比较
			每一轮比较后，最大的元素会被移动到最后的位置，因此后续的比较就不需要再考虑这个元素了。
			如果当前元素大于下一个元素，则交换它们
*/

// bubbleSort 升序
func bubbleSort() {
	BubbleSort([]int{3, 2, 1, 5, 4})
}

func BubbleSort(arr []int) []int {
	n := len(arr)
	// 外层循环控制需要进行多少轮比较
	for i := 0; i < n; i++ {
		// 每一轮比较后，最大的元素会被移动到最后的位置，因此后续的比较就不需要再考虑这个元素了
		for j := 0; j < n-i-1; j++ {
			// 如果当前元素大于下一个元素，则交换它们
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
