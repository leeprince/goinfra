package skilldoublepointer

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/5 23:43
 * @Desc:	合并两个有序链表
 *				代码中还用到一个链表的算法题中是很常见的「虚拟头结点」技巧，也就是 newListNode 节点。你可以试试，如果不使用 newListNode 虚拟节点，代码会复杂一些，需要额外处理指针 p 为空的情况。而有了 dummy 节点这个占位符，可以避免处理空指针的情况，降低代码的复杂性。
 *				- 什么时候需要用虚拟头结点？
 *					总结下：当你需要创造一条新链表的时候，可以使用虚拟头结点简化边界情况的处理。比如说，让你把两条有序链表合并成一条新的有序链表，是不是要创造一条新链表？再比你想把一条链表分解成两条链表，是不是也在创造新链表？这些情况都可以使用虚拟头结点简化边界情况的处理。
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
	MergeTwoListsV1(&ListNode{}, &ListNode{})
}

func MergeTwoListsV1(list1 *ListNode, list2 *ListNode) *ListNode {
	// 虚拟头节点
	newListNode := &ListNode{
		Val:  -1,
		Next: nil,
	}

	// 保持原始链表
	p := newListNode
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	// 易错点：必须使用原有的newListNode变量非p变量，因为p变量的Next指针是在不断变化的
	newListNode = newListNode.Next

	return newListNode
}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

func MergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟头结点
	dummy := &ListNode{-1, nil}
	p := dummy
	p1 := l1
	p2 := l2

	for p1 != nil && p2 != nil {
		// 比较 p1 和 p2 两个指针
		// 将值较小的的节点接到 p 指针
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		// p 指针不断前进
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	dummy = dummy.Next
	return dummy
}
