package arrayslice

import (
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/26 22:40
 * @Desc:
 */

func TestInsertSortV1(t *testing.T) {
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
				arr: []int{5, 5},
			},
			want: []int{5, 5},
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 4, 3},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				arr: []int{5, 2, 5, 3, 5, 1},
			},
			want: []int{1, 2, 3, 5, 5, 5},
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
			if got := InsertSortV1(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSortV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertSortV2(t *testing.T) {
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
				arr: []int{5, 5},
			},
			want: []int{5, 5},
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 4, 3},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				arr: []int{5, 2, 5, 3, 5, 1},
			},
			want: []int{1, 2, 3, 5, 5, 5},
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
			if got := InsertSortV2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSortV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
