package arrayslice

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/9 01:23
 * @Desc:
 */

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 3, 3, 4},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 3, 3, 4, 4},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 3, 3, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatesV2(tt.args.nums); got != tt.want {
				t.Errorf("RemoveDuplicatesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
