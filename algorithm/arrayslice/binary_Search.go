package arrayslice

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/27 01:48
 * @Desc:	二分法查找:在已排序好的数组上查找匹配元素的索引，找不到返回-1
 * 				1. 确定该区间的中间位置K
 *				2. 将查找的值T与array[k]比较。若相等，查找成功返回此位置；否则确定新的查找区域，继续二分查找。
 *				3. 区域确定如下：
 *					1) a.array[k]>T 由数组的有序性可知array[k,k+1,……,high]>T;故新的区间为array[low,……，K-1]
 *					2) b.array[k]<T 类似上面查找区间为array[k+1,……，high]。
 *					3) 每一次查找与中间值比较，可以确定是否查找成功，不成功当前查找区间将缩小一半，递归查找即可。时间复杂度为:O(log2n)。
 */

func BinarySearch() {
	arr := []int{1, 3, 5, 7, 9}
	v := 2
	BinarySearchV1(arr, v)
}

func BinarySearchV1(arr []int, v int) int {
	left := 0
	right := len(arr) - 1
	
	for left <= right {
		midIndex := (right-left)/2 + left // 推荐。比 midIndex := (right + left) / 2 写法更易读
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
