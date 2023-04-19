package algorithm

import (
	"github.com/leeprince/goinfra/algorithm/skilldoublepointer"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/5 23:30
 * @Desc:	双指针技巧
 *			单链表有很多巧妙的操作，本文就总结一下单链表的基本技巧，每个技巧都对应着至少一道算法题：
 *				1、合并两个有序链表
 *				2、链表的分解
 *				3、合并 k 个有序链表
 *				4、寻找单链表的倒数第 k 个节点
 *				5、寻找单链表的中点
 *				6、判断单链表是否包含环并找出环起点
 *				7、判断两个单链表是否相交并找出交点
 *			这些解法都用到了双指针技巧，所以说对于单链表相关的题目，双指针的运用是非常广泛的
 */

func DoublePointer() {
	skilldoublepointer.MergeTwoLists() // 合并两个有序链表
	skilldoublepointer.Partition()     // 分隔链表

}
