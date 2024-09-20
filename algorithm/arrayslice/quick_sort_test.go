package arrayslice

import (
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/21 17:26
 * @Desc:
 */

func TestQuickSortV1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				arr: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				arr: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			args: args{
				arr: []int{3, 2, 1, 4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				arr: []int{5, 2, 3, 4, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "",
			args: args{
				arr: []int{5, 2, 5, 3, 4, 1},
			},
			want: []int{1, 2, 3, 4, 5, 5},
		},
		{
			name: "",
			args: args{
				arr: []int{95, 45, 15, 78, 84, 51, 24, 12},
			},
			want: []int{12, 15, 24, 45, 51, 78, 84, 95},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
