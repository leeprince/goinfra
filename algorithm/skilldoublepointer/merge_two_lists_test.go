package skilldoublepointer

import (
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/6 00:30
 * @Desc:
 */

func TestMergeTwoListsV1(t *testing.T) {
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
		// 判断结果与期望值是否一致方法一【推荐：简单、原生】
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoListsV1(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf(">1: MergeTwoListsV1() = %v, want %v", got, tt.want)
			}
		})
		// // 判断结果与期望值是否一致方法二
		// t.Run(tt.name, func(t *testing.T) {
		// 	got := MergeTwoListsV1(tt.args.list1, tt.args.list2)
		// 	// got := MergeTwoListsV2(tt.args.list1, tt.args.list2)
		// 	want := tt.want
		// 	for got != nil && want != nil {
		// 		if got.Val != want.Val {
		// 			t.Errorf(">2: MergeTwoListsV1() = %v, want %v", got, want)
		// 		}
		// 		got = got.Next
		// 		want = want.Next
		// 	}
		// })
	}
}
