package characterutil

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/8 17:27
 * @Desc:
 */

func TestIsUTF8(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				v: "dfa",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				v: "\xb2\xe2\xca\xd4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUTF8(tt.args.v); got != tt.want {
				t.Errorf("IsUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}
