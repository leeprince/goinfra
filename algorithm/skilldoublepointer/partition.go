package skilldoublepointer

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/11 01:35
 * @Desc:	分隔链表: 注意虚拟头结点的运用
 */

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。



示例 1：

输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

示例 2：

输入：head = [2,1], x = 2
输出：[1,2]



提示：

    链表中节点的数目在范围 [0, 200] 内
    -100 <= Node.val <= 100
    -200 <= x <= 200

https://leetcode.cn/problems/partition-list/
*/

func Partition() {
	PartitionV1(&ListNode{}, 3)
}
func PartitionV1(head *ListNode, x int) *ListNode {
	// 左链表
	left := &ListNode{
		Val:  -1,
		Next: nil,
	}
	// 右链表
	right := &ListNode{
		Val:  -1,
		Next: nil,
	}
	// 复制可移动指针的左右链表
	l, r := left, right

	// 复制可移动指针的初始链表
	h := head

	// 遍历初始链表
	for h != nil {
		// 左右链表赋值，并初始化
		if h.Val >= x {
			r.Next = h
			r = r.Next
		} else {
			l.Next = h
			l = l.Next
		}
		// 移动初始链表，为避免移动的初始链表影响上面链表的指针，赋值到临时变量再赋值
		tmp := h.Next
		h.Next = nil
		h = tmp
	}
	// 合并左右链表: 左链表指向去除虚拟头节点后的右链表
	l.Next = right.Next

	// 新链表指向去除虚拟头节点后的左链表
	newListNode := left.Next

	return newListNode
}
