package arrayslice

// 删除排序数组中的重复项
/*
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。

将最终结果插入 nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

判题标准:

系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = RemoveDuplicatesV1(nums); // 调用

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}

如果所有断言都通过，那么您的题解将被 通过。



示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。



提示：

    1 <= nums.length <= 3 * 104
    -104 <= nums[i] <= 104
    nums 已按 升序 排列

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2gy9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func RemoveDuplicates() (respInt int) {
	slice := []int{1, 2, 3, 4}

	/*
		解题思路：快慢双指针
		1. 初始化慢指针为0，快指针为1
		2. 快慢指针指向的值进行比较
		3. 快慢指针指向的值一样时：快指针右移1位；慢指针不动
		4. 快慢指针指向的值不一样时：慢指针+1索引的值与快指针索引的值交换；快指针右移1位；慢指针右移1位
	*/
	/*
		待优化：
		1. 内存优化
			1) 返回的新数组长度，通过慢指针统计
			2) 仅处理不一样的情况，否则继续循环
	*/
	respInt = RemoveDuplicatesV1(slice)

	/*
		解题思路：快慢双指针
		1. 初始化慢指针为1，快指针为1
		2. 快指针与快指针-1指向的值比较
		3. 快慢指针指向的值一样时：快指针右移1位；慢指针不动
		4. 快慢指针指向的值不一样时：慢指针的值等于快指针的值，快指针指向的值无需处理；快指针右移1位；慢指针右移1位
	*/
	respInt = RemoveDuplicatesV2(slice)

	return

}

func RemoveDuplicatesV1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 0
	respInt := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			temp := nums[slow+1]
			nums[slow+1] = nums[fast]
			nums[fast] = temp
			slow++
			respInt++
		}
	}

	return respInt
}

func RemoveDuplicatesV2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
