package arrayslice

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/21 01:43
 * @Desc:	快速排序
 *				1. 首先设定一个分界值，通过该分界值将数组分成左右两部分
 *				2. 将大于或等于分界值的数据集中到数组右边，小于分界值的数据集中到数组的左边。此时，左边部分中各元素都小于分界值，而右边部分中各元素都大于或等于分界值。
 *				3. 然后，左边和右边的数据可以独立排序。对于左侧的数组数据，又可以取一个分界值，将该部分数据分成左右两部分，同样在左边放置较小值，右边放置较大值。右侧的数组数据也可以做类似处理
 *				4. 重复上述过程，可以看出，这是一个递归定义。通过递归将左侧部分排好序后，再递归排好右侧部分的顺序。当左、右两个部分各数据排序完成后，整个数组的排序也就完成了。
 */

func quickSort() {
	QuickSort([]int{5, 2, 1, 3, 4})
}

// 升序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 选择基准值，这里选择第一个元素
	midd := arr[0]
	
	var left []int
	var right []int
	
	for i := 1; i < len(arr); i++ {
		if midd < arr[i] {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}
	
	left = QuickSort(left)
	right = QuickSort(right)
	
	resp := append(left, midd)
	resp = append(resp, right...)
	return resp
}
