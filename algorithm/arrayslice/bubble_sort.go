package arrayslice

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/21 00:19
 * @Desc:	冒泡排序
 *				1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
 *				2. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
 *				3. 针对所有的元素重复以上的步骤，除了最后一个。
 *				4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较
 */

func BubbleSort() {
	req := []int64{5, 2, 1, 3, 4}
	BubbleSortV1(req)
}

// 升序
func BubbleSortV1(arr []int64) []int64 {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
