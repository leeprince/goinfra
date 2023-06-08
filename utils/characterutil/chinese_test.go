package characterutil

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/8 17:24
 * @Desc:
 */

func TestIsChinese(t *testing.T) {
	type args struct {
		v []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				v: []rune("我爱中国"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				v: []rune("I Love You!"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChinese(tt.args.v); got != tt.want {
				t.Errorf("IsChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}
