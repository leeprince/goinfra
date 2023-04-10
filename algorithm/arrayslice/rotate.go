package arrayslice

import "github.com/leeprince/goinfra/algorithm/pkg"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/13 20:23
 * @Desc:	旋转数组
 */

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。



示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]



提示：

    1 <= nums.length <= 105
    -231 <= nums[i] <= 231 - 1
    0 <= k <= 105



进阶：

    尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
    你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func Rotate() {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 3

	/*
		遍历移动到额外数组
		1. 状态：
			移动位：k%len(nums)
		2. 状态转移：
			遍历nums
			numsNew[i+k%len(nums)]=nums[i]
		3. 初始值：
			遍历nums从索引0开始
		4. 输出值：
			nums=numsNew

		使用额外数组：空间复杂度O(n)
	*/
	RotateV1(nums, k)

	/*
		翻转代替旋转
		1. 定义状态
			翻转：nums[len(nums)-1-i],nums[i]=nums[i],nums[len(nums)-1-i]
		2. 状态转移
			全部翻转：nums[len(nums)-1-i],nums[i]=nums[i],nums[len(nums)-1-i]
			再部分翻转：[0, k%len(nums)-1],[k%len(nums), len(nums)-1]
		3. 初始值
			整个nums
		4. 输出值
			部分翻转后的结果前后组合在一起

		空间复杂度O(1)
	*/
	RotateV2(nums, k)
}

func RotateV1(nums []int, k int) {
	move := k % len(nums)
	if move == 0 {
		return
	}
	numsNew := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// 到切片取值到新切片
		numsNew[(i+move)%len(nums)] = nums[i]
	}

	// 错误的赋值方式(nums = numsNew)，应使用copy将元素从源切片复制到目标切片或者遍历赋值
	copy(nums, numsNew)
	// for i, i2 := range numsNew {
	// 	nums[i] = i2
	// }
}

func RotateV2(nums []int, k int) {
	l := len(nums)
	move := k % l
	if move == 0 {
		return
	}

	// 全部翻转
	pkg.ReverseSlice(nums)

	// 部分翻转：切片左闭右开
	pkg.ReverseSlice(nums[:move])
	pkg.ReverseSlice(nums[move:])
}
