package utils

import (
    "testing"
    "time"
)

/**
 * @Author: prince.lee
 * @Date:   2022/3/24 17:13
 * @Desc:
 */

func TestUseMillisecondUnit(t *testing.T) {
    type args struct {
        dur time.Duration
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            args: args{
                dur: 0,
            },
            want: true,
        },
        {
            args: args{
                dur: -1,
            },
            want: true,
        },
        {
            args: args{
                dur: time.Millisecond * 100,
            },
            want: true,
        },
        {
            args: args{
                dur: time.Millisecond * 1000,
            },
        },
        {
            args: args{
                dur: time.Millisecond * 2000,
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := UseMillisecondUnit(tt.args.dur); got != tt.want {
                t.Errorf("UseMillisecondUnit() = %v, want %v", got, tt.want)
            }
        })
    }
}