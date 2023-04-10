package arrayslice

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/13 23:39
 * @Desc:
 */

func Test_rotateV1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		args     args
		wantResp []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			wantResp: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			wantResp: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateV1(tt.args.nums, tt.args.k)
			for i := 0; i < len(tt.args.nums); i++ {
				if tt.wantResp[i] != tt.args.nums[i] {
					t.Errorf("RotateV1() = %v, want %v", tt.args.nums, tt.wantResp)
					return
				}
			}
		})
	}
}

func Test_rotateV2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		args     args
		wantResp []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			wantResp: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			wantResp: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateV2(tt.args.nums, tt.args.k)
			for i := 0; i < len(tt.args.nums); i++ {
				if tt.wantResp[i] != tt.args.nums[i] {
					t.Errorf("RotateV2() = %v, want %v", tt.args.nums, tt.wantResp)
					return
				}
			}
		})
	}
}
