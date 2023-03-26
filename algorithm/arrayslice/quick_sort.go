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

func QuickSort() {
	req := []int64{5, 2, 1, 3, 4}

	// 需要创建较多切片
	quickSortV1(req)

	// 无需创建较多切片：结合左右指针
	quickSortV2(req)
}

// 升序
func quickSortV1(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}
	var left []int64
	var right []int64
	var resp []int64

	midd := arr[0]

	for i := 1; i < len(arr); i++ {
		if midd < arr[i] {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	left = quickSortV1(left)
	right = quickSortV1(right)

	resp = append(left, midd)
	resp = append(resp, right...)
	return resp
}

func quickSortV2(arr []int64) []int64 {
	quickSortV21(arr, 0, len(arr)-1)
	return arr
}

func quickSortV21(arr []int64, left, right int) {
	if left >= right {
		return
	}

	midd := arr[left]
	i, j := left, right

	for left < right {
		if arr[left] >= midd && arr[right] <= midd {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
			continue
		}
		if arr[left] <= midd {
			left++
		}
		if arr[right] >= midd {
			right--
		}
	}

	// 易漏点：出现这个情况的原因是：左右指针相邻，且需要满足`arr[left] >= midd && arr[right] <= midd`时
	if left > right {
		left--
		right++
	}

	quickSortV21(arr, i, left)
	quickSortV21(arr, right, j)
}
