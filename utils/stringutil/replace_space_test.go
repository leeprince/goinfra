package stringutil

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/14 16:42
 * @Desc:
 */

func TestReplaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpace(tt.args.s); got != tt.want {
				t.Errorf("ReplaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
