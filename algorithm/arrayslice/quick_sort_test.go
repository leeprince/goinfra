package arrayslice

import (
	"reflect"
	"sort"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/21 17:26
 * @Desc:
 */

func TestQuickSortV1(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "",
			args: args{
				arr: []int64{1},
			},
			want: []int64{1},
		},
		{
			name: "",
			args: args{
				arr: []int64{3, 2, 1},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "",
			args: args{
				arr: []int64{3, 2, 1, 4},
			},
			want: []int64{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				arr: []int64{5, 2, 3, 4, 1},
			},
			want: []int64{1, 2, 3, 4, 5},
		},
		{
			name: "",
			args: args{
				arr: []int64{5, 2, 5, 3, 4, 1},
			},
			want: []int64{1, 2, 3, 4, 5, 5},
		},
		{
			name: "",
			args: args{
				arr: []int64{95, 45, 15, 78, 84, 51, 24, 12},
			},
			want: []int64{12, 15, 24, 45, 51, 78, 84, 95},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortV1(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortV11(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				arr: []int64{1},
			},
		},
		{
			name: "",
			args: args{
				arr: []int64{3, 2, 1},
			},
		},
		{
			name: "",
			args: args{
				arr: []int64{3, 2, 1, 4},
			},
		},
		{
			name: "",
			args: args{
				arr: []int64{5, 2, 3, 4, 1},
			},
		},
		{
			name: "",
			args: args{
				arr: []int64{5, 2, 5, 3, 4, 1},
			},
		},
		{
			name: "",
			args: args{
				arr: []int64{95, 45, 15, 78, 84, 51, 24, 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickSortV1(tt.args.arr)
			if !sort.SliceIsSorted(got, func(i, j int) bool {
				if i < j {
					return true
				}
				return false
			}) {
				t.Errorf("QuickSortV1() = %v, tt.args.arr %v", got, tt.args.arr)
			}
		})
	}
}

func TestQuickSortV2(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "",
			args: args{
				arr: []int64{1},
			},
			want: []int64{1},
		},
		{
			name: "",
			args: args{
				arr: []int64{3, 2, 1},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "",
			args: args{
				arr: []int64{3, 2, 1, 4},
			},
			want: []int64{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				arr: []int64{5, 2, 3, 4, 1},
			},
			want: []int64{1, 2, 3, 4, 5},
		},
		{
			name: "",
			args: args{
				arr: []int64{5, 2, 5, 3, 4, 1},
			},
			want: []int64{1, 2, 3, 4, 5, 5},
		},
		{
			name: "",
			args: args{
				arr: []int64{5, 5},
			},
			want: []int64{5, 5},
		},
		{
			name: "",
			args: args{
				arr: []int64{1, 2, 4, 3},
			},
			want: []int64{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				arr: []int64{5, 2, 5, 3, 5, 1},
			},
			want: []int64{1, 2, 3, 5, 5, 5},
		},
		{
			name: "",
			args: args{
				arr: []int64{95, 45, 15, 78, 84, 51, 24, 12},
			},
			want: []int64{12, 15, 24, 45, 51, 78, 84, 95},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortV2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
