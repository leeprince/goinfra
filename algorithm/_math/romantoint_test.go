package _math

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/2/23 14:37
 * @Desc:
 */

func Test_romanToIntPractice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		wantResp int
	}{
		{
			name: "",
			args: args{
				s: "III",
			},
			wantResp: 3,
		},
		{
			name: "",
			args: args{
				s: "XLIX",
			},
			wantResp: 49,
		},
		{
			name: "",
			args: args{
				s: "CMXCIX",
			},
			wantResp: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := romanToIntPractice1(tt.args.s); gotResp != tt.wantResp {
				t.Errorf("romanToIntPractice1() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResp := romanToIntBeat(tt.args.s); gotResp != tt.wantResp {
				t.Errorf("romanToIntPractice1() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
