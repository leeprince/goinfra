package stringutil

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/15 13:36
 * @Desc:
 */

func TestFillChar(t *testing.T) {
	type args struct {
		sourceStr  string
		fillChar   rune
		fillLen    int
		IsFillLeft bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				sourceStr:  "12345",
				fillChar:   'A',
				fillLen:    10,
				IsFillLeft: false,
			},
			want: "12345AAAAA",
		},
		{
			name: "",
			args: args{
				sourceStr:  "12345",
				fillChar:   'A',
				fillLen:    10,
				IsFillLeft: true,
			},
			want: "AAAAA12345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FillChar(tt.args.sourceStr, tt.args.fillChar, tt.args.fillLen, tt.args.IsFillLeft); got != tt.want {
				t.Errorf("FillChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
