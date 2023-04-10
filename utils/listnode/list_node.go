package listnode

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/8 23:12
 * @Desc:	链表
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func IsEqualListNode(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
