package test

import (
	"github.com/leeprince/goinfra/algorithm/arrayslice"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/21 00:38
 * @Desc:
 */

func TestBubbleSortV1(t *testing.T) {
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
			if got := arrayslice.BubbleSortV1(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSortV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSortV11(t *testing.T) {
	convey.Convey("冒泡排序测试", t, func() {
		convey.So(arrayslice.BubbleSortV1([]int64{
			5, 3, 1, 9, 7,
		}), convey.ShouldResemble, []int64{
			1, 3, 5, 7, 9,
		})
	})
}
