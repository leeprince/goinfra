package stringutil

import (
	"fmt"
	"testing"
)

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
		{
			name: "",
			args: args{
				s: "441521 20010116 XXXX    复制",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "    复    制    ",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "441521 20010116 XXXX    复制",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "   复  制  ",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpace(tt.args.s); got != tt.want {
				t.Errorf("ReplaceSpace() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("%s----\n", got)
			}
		})
	}
}

func TestReplaceWhitespaceChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "441521 20010116 XXXX    复制",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "    复    制    ",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "441521 20010116 XXXX    复制",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "   复  制  ",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceWhitespaceChar(tt.args.s); got != tt.want {
				t.Errorf("ReplaceSpace() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("%s----\n", got)
			}
		})
	}
}
