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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeKListsV1(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeKListsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
