package main

import (
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/9/23 23:38
 * @Desc:
 */

func TestSliceStr(t *testing.T) {
	type args struct {
		str string
		seq string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "SliceStr1",
			args: args{
				str: "a:b:c",
				seq: ":",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "SliceStr2",
			args: args{
				str: "",
				seq: ":",
			},
			want: []string{},
		},
		{
			name: "SliceStr3",
			args: args{
				str: "a:b:c",
				seq: "",
			},
			want: []string{"a:b:c"},
		},
		{
			name: "SliceStr4",
			args: args{
				str: "a:b:c:",
				seq: ":",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "SliceStr5",
			args: args{
				str: "abc",
				seq: ":",
			},
			want: []string{"abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceStr(tt.args.str, tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSliceStr(b *testing.B) {
	// time.Sleep(time.Second * 2)
	b.ResetTimer() // 重置计时器，避免每次都从开始计时，影响测试结果。可以在连接数据库、Redis等耗时操作之后调用
	for i := 0; i < b.N; i++ {
		SliceStr("a:b:c", ":")
	}
}
