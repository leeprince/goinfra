package timeutil

import (
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/29 15:33
 * @Desc:
 */

func TestAfterSecond(t *testing.T) {
	type args struct {
		second int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "",
			args: args{
				second: 10,
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(time.Now().Format(consts.TimeYmdHis))

			got := AfterSecond(tt.args.second)

			fmt.Println("AfterSecond:", got.Format(consts.TimeYmdHis))
		})
	}
}
