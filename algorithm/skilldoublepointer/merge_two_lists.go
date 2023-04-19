package skilldoublepointer

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

/*
形象地理解，这个算法的逻辑类似于拉拉链，l1, l2 类似于拉链两侧的锯齿，指针 p 就好像拉链的拉索，将两个有序链表合并；或者说这个过程像蛋白酶合成蛋白质，l1, l2 就好比两条氨基酸，而指针 p 就好像蛋白酶，将氨基酸组合成蛋白质。

代码中还用到一个链表的算法题中是很常见的「虚拟头结点」技巧，也就是 newListNode 节点。你可以试试，如果不使用 newListNode 虚拟节点，代码会复杂一些，需要额外处理指针 p 为空的情况。而有了 dummy 节点这个占位符，可以避免处理空指针的情况，降低代码的复杂性。
	- 什么时候需要用虚拟头结点？
	总结下：当你需要创造一条新链表的时候，可以使用虚拟头结点简化边界情况的处理。比如说，让你把两条有序链表合并成一条新的有序链表，是不是要创造一条新链表？再比你想把一条链表分解成两条链表，是不是也在创造新链表？这些情况都可以使用虚拟头结点简化边界情况的处理。
*/

func MergeTwoLists() {
	MergeTwoListsV1(&ListNode{}, &ListNode{})
}

func MergeTwoListsV1(list1 *ListNode, list2 *ListNode) *ListNode {
	// 虚拟头节点
	newListNode := &ListNode{
		Val:  -1,
		Next: nil,
	}

	// 保持原始链表：复制为可移动指针的新链表
	p := newListNode
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		// 比较 p1 和 p2 两个指针
		// 将值较小的的节点接到 p 指针
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
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

	// 易错点：必须使用原有的newListNode变量非p变量去除虚拟节点，因为p变量的Next指针是在不断变化的。
	newListNode = newListNode.Next

	return newListNode
}
