package arrayslice

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/27 20:08
 * @Desc:
 */

func Test_binarySearchV1(t *testing.T) {
	type args struct {
		arr []int
		v   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{5},
				v:   5,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				v:   1,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				v:   3,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				v:   5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				v:   7,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				v:   9,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchV1(tt.args.arr, tt.args.v); got != tt.want {
				t.Errorf("binarySearchV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
