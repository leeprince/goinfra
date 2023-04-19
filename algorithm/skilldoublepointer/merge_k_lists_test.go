package skilldoublepointer

import (
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/19 10:25
 * @Desc:
 */

func TestMergeKListsV1(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				lists: []*ListNode{
					{
						Val: 1,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
					{
						Val: 2,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  8,
								Next: nil,
							},
						},
					},
					{
						Val: 3,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val:  9,
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
											Val: 8,
											Next: &ListNode{
												Val:  9,
												Next: nil,
											},
										},
									},
								},
							}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeKListsV1(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeKListsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
