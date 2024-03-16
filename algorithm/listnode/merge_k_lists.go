package listnode

import "container/heap"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/19 09:40
 * @Desc:	合并 K 个升序链表
 */

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：

输入：lists = []
输出：[]

示例 3：

输入：lists = [[]]
输出：[]



提示：

    k == lists.length
    0 <= k <= 10^4
    0 <= lists[i].length <= 500
    -10^4 <= lists[i][j] <= 10^4
    lists[i] 按 升序 排列
    lists[i].length 的总和不超过 10^4

https://leetcode.cn/problems/merge-k-sorted-lists/
*/

/*
合并 k 个有序链表的逻辑类似合并两个有序链表，难点在于，如何快速得到 k 个节点中的最小节点，接到结果链表上？

这里我们就要用到 优先级队列（二叉堆） 这种数据结构，把链表节点放入一个最小堆，就可以每次获得 k 个节点中的最小节点：

优先队列 pq 中的元素个数最多是 k，所以一次 poll 或者 add 方法的时间复杂度是 O(logk)；所有的链表节点都会被加入和弹出 pq，所以算法整体的时间复杂度是 O(Nlogk)，其中 k 是链表的条数，N 是这些链表的节点总数。
*/

func MergeKLists() {
	MergeKListsV1([]*ListNode{})
}

func MergeKListsV1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	
	newList := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dumpNewList := newList
	
	minHeap := make(MinHeap, 0)
	heap.Init(&minHeap)
	
	for _, list := range lists {
		if list != nil {
			heap.Push(&minHeap, list)
		}
	}
	
	for minHeap.Len() > 0 {
		minNode := heap.Pop(&minHeap).(*ListNode)
		dumpNewList.Next = minNode
		if minNode.Next != nil {
			heap.Push(&minHeap, minNode.Next)
		}
		dumpNewList = dumpNewList.Next
	}
	
	// 取出虚拟节点
	newList = newList.Next
	
	return newList
}

// 小根堆
type MinHeap []*ListNode

func (p MinHeap) Len() int {
	return len(p)
}

func (p MinHeap) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p MinHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *MinHeap) Push(x any) {
	node := x.(*ListNode)
	*p = append(*p, node)
}

func (p *MinHeap) Pop() any {
	l := len(*p)
	pDump := *p
	node := pDump[l-1]
	*p = pDump[:l-1]
	return node
}
