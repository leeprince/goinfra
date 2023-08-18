package timeutil

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/15 17:00
 * @Desc:
 */

func TestCheckDateFormat(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				date: "202308",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				date: "202307",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				date: "202318",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				date: "102312",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				date: "102322",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckDateFormat(tt.args.date); got != tt.want {
				t.Errorf("CheckDateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
