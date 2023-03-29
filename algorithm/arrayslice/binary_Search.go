package arrayslice

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/27 01:48
 * @Desc:	二分法查找:在已排序好的数组上查找匹配元素的索引，找不到返回-1
 */

func BinarySearch() {
	arr := []int{1, 3, 5, 7, 9}
	v := 2
	binarySearchV1(arr, v)
}

func binarySearchV1(arr []int, v int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		midIndex := (right-left)/2 + left
		midNum := arr[midIndex]
		if midNum == v {
			return midIndex
		}
		if midNum > v {
			right = midIndex - 1
		} else {
			left = midIndex + 1
		}
	}

	return -1
}

// 生成二分法查找
