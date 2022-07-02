package utils

import (
    "fmt"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/28 下午11:24
 * @Desc:
 */

const (
    strEN = "leeprince"
    strZH = "百炼成钢"
)

func TestGbkToUtf8AndUtf8ToGbk(t *testing.T) {
    type args struct {
        s []byte
    }
    tests := []struct {
        name    string
        args    args
        want    []byte
        wantErr bool
    }{
        {
            args:args{
                s: []byte(strEN),
            },
        },
        {
            args:args{
                s: []byte(strZH),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gbktoutf8, err := GbkToUtf8(tt.args.s)
            if (err != nil) != tt.wantErr {
                t.Errorf("GbkToUtf8() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            fmt.Println("GbkToUtf8:", string(gbktoutf8))
            
            utf8togbk, err := Utf8ToGbk(gbktoutf8)
            if (err != nil) != tt.wantErr {
                t.Errorf("GbkToUtf8() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            fmt.Println("Utf8ToGbk:", string(utf8togbk))
        })
    }
}