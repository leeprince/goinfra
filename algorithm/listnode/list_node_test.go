package listnode

import (
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/5 23:30
 * @Desc:	双指针技巧
 *			单链表有很多巧妙的操作，本文就总结一下单链表的基本技巧，每个技巧都对应着至少一道算法题：
 *				1、合并两个有序链表
 *				2、链表的分解
 *				3、
 *				4、寻找单链表的倒数第 k 个节点
 *				5、寻找单链表的中点
 *				6、判断单链表是否包含环并找出环起点
 *				7、判断两个单链表是否相交并找出交点
 *			这些解法都用到了双指针技巧，所以说对于单链表相关的题目，双指针的运用是非常广泛的
 */

// 合并两个有序链表
func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				list2: nil,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			name: "",
			args: args{
				list1: &ListNode{
					Val:  1,
					Next: nil,
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
		{
			name: "",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
				list2: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val:  8,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
									Next: &ListNode{
										Val: 7,
										Next: &ListNode{
											Val:  8,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		// reflect.DeepEqual 和 遍历链表都可以判断结果与期望值是否一致
		// 判断结果与期望值是否一致方法一
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf(">1: MergeTwoListsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 链表的分解
func TestPartition(t *testing.T) {
	Partition() // 链表的分解
}

// 合并 k 个有序链表
func TestMergeKLists(t *testing.T) {
	MergeKLists()
}
