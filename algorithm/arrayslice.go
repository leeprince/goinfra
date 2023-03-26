package algorithm

import "github.com/leeprince/goinfra/algorithm/arrayslice"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/9 01:38
 * @Desc:   数组
 */

func ArraySlice() {
	arrayslice.BubbleSort()       // 冒泡排序：两次遍历，一次找元素，一次遍历比较确定最大值：比较相邻值，小的在左大的在右，最后一个即为最大值。
	arrayslice.QuickSort()        // 快速排序：定分界值，找出小于等于分界值的左和大于分界值的右边，在分别对左右两边继续递归。可以结合左右指针优化
	arrayslice.HeapSort()         // 堆排序：构建大根堆，堆顶与堆尾替换，替换后数组-1并重新在堆顶构建大根堆
	arrayslice.InsertSort()       // 插入排序：从1个元素开始，从右边待排序序列的最左边元素开始在左边的有序序列中找到合适位置插入，且插入后的所有元素右移1位
	arrayslice.BinarySearch()     // 二分法查找 // TODO:  - prince@todo 2023/3/27 02:03
	arrayslice.RemoveDuplicates() // 删除排序数组中的重复项：快慢双指针
	arrayslice.MaxProfit()        // 买卖股票的最佳时机 II：回溯算法/动态规划/贪心算法
	arrayslice.Rotate()           // 旋转数组：翻转代替旋转
}
