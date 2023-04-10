package arrayslice

import "github.com/leeprince/goinfra/algorithm/pkg"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/26 21:12
 * @Desc:   插入排序
 *             1. 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
 *             2. 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
 */

func InsertSort() {
	arr := []int{5, 3, 1, 4, 2}

	// 代码有点多,且每次需要从前面开始遍历
	InsertSortV1(arr)

	// 优化：减少代码行,且减少遍历小于当前元素的前面序列
	InsertSortV2(arr)

	// 优化：遍历前面已排序的序列时，进行对半查找效率更高，不过会变为不稳定排序
	InsertSortV3(arr)
}

func InsertSortV1(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			// 是否需要向右移动
			var isNeedRight bool
			if !isNeedRight && arr[i] <= arr[j] {
				pkg.SwapInt(arr, i, j)
				isNeedRight = true
				continue
			}
			if isNeedRight {
				pkg.SwapInt(arr, j, j+1)
			}
		}
	}
	return arr
}

func InsertSortV2(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			// 右边大于左边就开始向右移动，直到等于或者小于时结束循环
			if arr[j] <= arr[j+1] {
				break
			}
			pkg.SwapInt(arr, j, j+1)
		}
	}
	return arr
}

func InsertSortV3(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// 对半查找，要么匹配到其中一个，要么在两个值中间去左边位置+1，且后面所有元素右移
		// TODO:  - prince@todo 2023/3/27 02:02
	}
	return arr
}
