package doublepointer

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/5 23:43
 * @Desc:	合并两个有序链表
 */

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：

输入：l1 = [], l2 = []
输出：[]

示例 3：

输入：l1 = [], l2 = [0]
输出：[0]



提示：

    两个链表的节点数目范围是 [0, 50]
    -100 <= Node.val <= 100
    l1 和 l2 均按 非递减顺序 排列

https://leetcode.cn/problems/merge-two-sorted-lists/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists() {
	mergeTwoLists(&ListNode{}, &ListNode{})
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    p1, p2 := 0, 0
    newListNode := &ListNode{
        Val = 0
        Next = &ListNode{}
    }

    for list1.Next != nil && lsit2.ListNode != nil {
        if list1.Val < list2.Val {
            newListNode.Next.Val = list1.Val
            newListNode.Next.Next = &ListNode{}
            p1++
        } else {
            newListNode.Next.Val = list2.Val
            newListNode.Next.Next = &ListNode{}
            p1++
        }
    }

	return nil
}



















