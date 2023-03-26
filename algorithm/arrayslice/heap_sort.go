package arrayslice

import "github.com/leeprince/goinfra/algorithm/pkg"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/26 12:33
 * @Desc:	堆排序:堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。堆排序可以说是一种利用堆的概念来排序的选择排序。分为两种方法：
 *		    大根堆：每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列；
 *		    小根堆：每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列；
 *          // 索引从0开始。arr[0]代表根节点
 *		    // 父节点索引找子节点索引：父节点i,左子节点：i*2+1;右子节点：i*2+2
 *			// 子节点索引找父节点索引：左子节点i,父节点：(i-1)/2;右子节点i,父节点：(i-2)/2;
 *			// 非叶子节点个数：len(arr)/2
 *			// 最后一个非叶子节点索引：(len(arr)/2)-1
 *
 *		    排序操作步骤：
 *              1. 构建一个大根堆 H[0……n-1]
 *              2. 把堆首（最大值）和堆尾互换
 *              3. 把堆的尺寸缩小 1，并重新构建堆（目的是把交换后的根节点数据调整到相应位置）
 *              4. 重复步骤 2，直到堆的尺寸为 1
 */

func HeapSort() {
	arr := []int{5, 3, 2, 1, 4}
	heapSortV1(arr)
}

func heapSortV1(arr []int) []int {
	l := len(arr)
	// 获取节点，在从下往上每个节点上构建大根堆，最终得到完整的大根堆
	for i := l / 2; i >= 0; i-- {
		heapMaxRoot(arr, i, l)
	}

	for i := l - 1; i >= 0; i-- {
		// 把堆首（最大值）和堆尾互换
		pkg.SwapInt(arr, 0, i)

		// 重新构建大根堆
		l--
		heapMaxRoot(arr, 0, l)
	}

	return arr
}

// 构建大根堆
func heapMaxRoot(arr []int, i, l int) {
	// i为父节点，确定左右节点索引
	left := i*2 + 1
	right := i*2 + 2
	// 最大值的索引，初始值为i
	maxNumIndex := i
	// 确定存在最大值的索引是否在左右节点
	if left < l && arr[left] > arr[maxNumIndex] {
		maxNumIndex = left
	}
	if right < l && arr[right] > arr[maxNumIndex] {
		maxNumIndex = right
	}
	// 不为当前根节点时，交换数据，并在当时最大值的节点上继续构建大根堆
	if maxNumIndex != i {
		pkg.SwapInt(arr, i, maxNumIndex)
		heapMaxRoot(arr, maxNumIndex, l)
	}
}
