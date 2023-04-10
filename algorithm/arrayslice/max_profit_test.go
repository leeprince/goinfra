package arrayslice

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/11 00:43
 * @Desc:
 */

func TestMaxProfitV1(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name     string
		args     args
		wantResp int
	}{
		{
			name: "#name00",
			args: args{
				prices: []int{1},
			},
			wantResp: 0,
		},
		{
			name: "#name01",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			wantResp: 7,
		},
		{
			name: "#name02",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			wantResp: 4,
		},
		{
			name: "#name03",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			wantResp: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 因为输出使用全局变量，所以maxProfitV1方法中应该初始化一次
			if gotResp := MaxProfitV1(tt.args.prices); gotResp != tt.wantResp {
				t.Errorf("MaxProfitV1() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestMaxProfitV2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name     string
		args     args
		wantResp int
	}{
		{
			name: "#name00",
			args: args{
				prices: []int{1},
			},
			wantResp: 0,
		},
		{
			name: "#name01",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			wantResp: 7,
		},
		{
			name: "#name02",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			wantResp: 4,
		},
		{
			name: "#name03",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			wantResp: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := MaxProfitV2(tt.args.prices); gotResp != tt.wantResp {
				t.Errorf("MaxProfitV2() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestMaxProfitV3(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name     string
		args     args
		wantResp int
	}{
		{
			name: "#name00",
			args: args{
				prices: []int{1},
			},
			wantResp: 0,
		},
		{
			name: "#name01",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			wantResp: 7,
		},
		{
			name: "#name02",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			wantResp: 4,
		},
		{
			name: "#name03",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			wantResp: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := MaxProfitV3(tt.args.prices); gotResp != tt.wantResp {
				t.Errorf("MaxProfitV3() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestMaxProfitV4(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name     string
		args     args
		wantResp int
	}{
		{
			name: "#name00",
			args: args{
				prices: []int{1},
			},
			wantResp: 0,
		},
		{
			name: "#name01",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			wantResp: 7,
		},
		{
			name: "#name02",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			wantResp: 4,
		},
		{
			name: "#name03",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			wantResp: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := MaxProfitV4(tt.args.prices); gotResp != tt.wantResp {
				t.Errorf("MaxProfitV4() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
