package arrayslice

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/9 01:38
 * @Desc:   数组/切片
 */

func ArraySlice() {
	BinarySearch() // 二分法查找：每次取中间位置M对应的值与查找值比较，查找值小于中间值，则目标值可能在左边[left,M-1],否则可能在右边[M+1, right]
	// ---
	bubbleSort() // 冒泡排序：相邻元素两两比较，大的往后移
	quickSort()  // 快速排序：选择基准元素，将数组分为两部分，递归排序
	heapSort()   // 堆排序：构建大根堆，堆顶与堆尾交换，缩小堆并重新调整
	InsertSort() // 插入排序：将待排序元素插入已排序序列的合适位置
	// ---
	RemoveDuplicates() // 删除排序数组中的重复项：快慢双指针
	MaxProfit()        // 买卖股票的最佳时机 II：回溯算法/动态规划/贪心算法
	Rotate()           // 旋转数组：翻转代替旋转
}
