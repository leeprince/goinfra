package contextutil

import (
	"context"
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/26 23:43
 * @Desc:
 */

func TestWriteHeaderContext(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *Header
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				data: &Header{
					UberTraceID: "",
					XRealIp:     "",
					LogId:       "LogId-01",
					Token:       "Token-01",
					AccessToken: "AccessToken-01",
				},
			},
			want: nil,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WriteHeaderContext(tt.args.ctx, tt.args.data)
			fmt.Printf("i:%d got:%+v \n", i, got)
			
			fmt.Println("--------")
			header, err := ReadHeaderByContext(got)
			fmt.Println(i, " ReadHeaderByContext err:", err)
			fmt.Printf("i:%d ReadHeaderByContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByContext header:", header)
			fmt.Printf("i:%d ReadHeaderByContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			header, err = ReadHeaderByCurrentContext(got)
			fmt.Println(i, " ReadHeaderByCurrentContext err:", err)
			fmt.Printf("i:%d ReadHeaderByCurrentContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByCurrentContext header:", header)
			fmt.Printf("i:%d ReadHeaderByCurrentContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			// 获取不到header值是正确的，因为在 WriteHeaderContext 是通过 metadata.AppendToOutgoingContext 添加 header 的，可以查看源码更清晰！！！
			header, err = ReadHeaderByIncomingContext(got)
			fmt.Println(i, " ReadHeaderByIncomingContext err:", err)
			fmt.Printf("i:%d ReadHeaderByIncomingContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByIncomingContext header:", header)
			fmt.Printf("i:%d ReadHeaderByIncomingContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			header, err = MustReadHeaderByIncomingContext(got)
			fmt.Println(i, " MustReadHeaderByIncomingContext err:", err)
			fmt.Printf("i:%d MustReadHeaderByIncomingContext err:%+v \n", i, err)
			fmt.Println(i, " MustReadHeaderByIncomingContext header:", header)
			fmt.Printf("i:%d MustReadHeaderByIncomingContext header:%+v \n", i, header)
		})
	}
}
