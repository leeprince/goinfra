package listnode

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/14 00:34
 * @Desc:
 */

// 定义单链表结构体
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一个节点的指针
}
