package contextutil

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/27 01:19
 * @Desc:
 */

func TestGetLogIdByContext(t *testing.T) {
	ctx := context.Background()
	
	type args struct {
		ctx *context.Context
	}
	tests := []struct {
		name      string
		args      args
		wantLogId string
	}{
		{
			name: "",
			args: args{
				ctx: &ctx,
			},
			wantLogId: "",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogId := LogIdByContext(tt.args.ctx)
			fmt.Println(i, "LogIdByContext:", gotLogId)
			
			fmt.Println("--------")
			header, err := ReadHeaderByContext(*tt.args.ctx)
			fmt.Println(i, " ReadHeaderByContext err:", err)
			fmt.Printf("i:%d ReadHeaderByContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByContext header:", header)
			fmt.Printf("i:%d ReadHeaderByContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			gotLogId = LogIdByContext(tt.args.ctx)
			fmt.Println(i, "LogIdByContext:", gotLogId)
			
			fmt.Println("--------")
			header, err = ReadHeaderByCurrentContext(*tt.args.ctx)
			fmt.Println(i, " ReadHeaderByCurrentContext err:", err)
			fmt.Printf("i:%d ReadHeaderByCurrentContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByCurrentContext header:", header)
			fmt.Printf("i:%d ReadHeaderByCurrentContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			// 获取不到header值是正确的，因为在 WriteHeaderContext 是通过 metadata.AppendToOutgoingContext 添加 header 的，可以查看源码更清晰！！！
			header, err = ReadHeaderByIncomingContext(*tt.args.ctx)
			fmt.Println(i, " ReadHeaderByIncomingContext err:", err)
			fmt.Printf("i:%d ReadHeaderByIncomingContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByIncomingContext header:", header)
			fmt.Printf("i:%d ReadHeaderByIncomingContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			header, err = MustReadHeaderByIncomingContext(*tt.args.ctx)
			fmt.Println(i, " MustReadHeaderByIncomingContext err:", err)
			fmt.Printf("i:%d MustReadHeaderByIncomingContext err:%+v \n", i, err)
			fmt.Println(i, " MustReadHeaderByIncomingContext header:", header)
			fmt.Printf("i:%d MustReadHeaderByIncomingContext header:%+v \n", i, header)
		})
	}
}

func TestGetLogIdByContext1(t *testing.T) {
	ctx := context.Background()
	
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		args      args
		wantLogId string
	}{
		{
			name: "",
			args: args{
				ctx: ctx,
			},
			wantLogId: "",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogId := LogIdByContext(&tt.args.ctx)
			fmt.Println(i, "LogIdByContext:", gotLogId)
			
			fmt.Println("--------")
			header, err := ReadHeaderByContext(tt.args.ctx)
			fmt.Println(i, " ReadHeaderByContext err:", err)
			fmt.Printf("i:%d ReadHeaderByContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByContext header:", header)
			fmt.Printf("i:%d ReadHeaderByContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			gotLogId = LogIdByContext(&tt.args.ctx)
			fmt.Println(i, "LogIdByContext:", gotLogId)
			
			fmt.Println("--------")
			header, err = ReadHeaderByCurrentContext(tt.args.ctx)
			fmt.Println(i, " ReadHeaderByCurrentContext err:", err)
			fmt.Printf("i:%d ReadHeaderByCurrentContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByCurrentContext header:", header)
			fmt.Printf("i:%d ReadHeaderByCurrentContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			// 获取不到header值是正确的，因为在 WriteHeaderContext 是通过 metadata.AppendToOutgoingContext 添加 header 的，可以查看源码更清晰！！！
			header, err = ReadHeaderByIncomingContext(tt.args.ctx)
			fmt.Println(i, " ReadHeaderByIncomingContext err:", err)
			fmt.Printf("i:%d ReadHeaderByIncomingContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByIncomingContext header:", header)
			fmt.Printf("i:%d ReadHeaderByIncomingContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			header, err = MustReadHeaderByIncomingContext(tt.args.ctx)
			fmt.Println(i, " MustReadHeaderByIncomingContext err:", err)
			fmt.Printf("i:%d MustReadHeaderByIncomingContext err:%+v \n", i, err)
			fmt.Println(i, " MustReadHeaderByIncomingContext header:", header)
			fmt.Printf("i:%d MustReadHeaderByIncomingContext header:%+v \n", i, header)
		})
	}
}

func TestLogIdByGinContext(t *testing.T) {
	ctx := gin.Context{}
	
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name      string
		args      args
		wantLogId string
	}{
		{
			name: "",
			args: args{
				ctx: &ctx,
			},
			wantLogId: "",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogId := LogIdByGinContext(tt.args.ctx)
			fmt.Println(i, "LogIdByGinContext:", gotLogId)
			
			fmt.Println("--------")
			header, err := ReadHeaderByContext(tt.args.ctx)
			fmt.Println(i, " ReadHeaderByContext err:", err)
			fmt.Printf("i:%d ReadHeaderByContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByContext header:", header)
			fmt.Printf("i:%d ReadHeaderByContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			gotLogId = LogIdByGinContext(tt.args.ctx)
			fmt.Println(i, "LogIdByGinContext:", gotLogId)
			
			fmt.Println("--------")
			header, err = ReadHeaderByCurrentContext(tt.args.ctx)
			fmt.Println(i, " ReadHeaderByCurrentContext err:", err)
			fmt.Printf("i:%d ReadHeaderByCurrentContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByCurrentContext header:", header)
			fmt.Printf("i:%d ReadHeaderByCurrentContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			// 获取不到header值是正确的，因为在 WriteHeaderContext 是通过 metadata.AppendToOutgoingContext 添加 header 的，可以查看源码更清晰！！！
			header, err = ReadHeaderByIncomingContext(tt.args.ctx)
			fmt.Println(i, " ReadHeaderByIncomingContext err:", err)
			fmt.Printf("i:%d ReadHeaderByIncomingContext err:%+v \n", i, err)
			fmt.Println(i, " ReadHeaderByIncomingContext header:", header)
			fmt.Printf("i:%d ReadHeaderByIncomingContext header:%+v \n", i, header)
			
			fmt.Println("--------")
			header, err = MustReadHeaderByIncomingContext(tt.args.ctx)
			fmt.Println(i, " MustReadHeaderByIncomingContext err:", err)
			fmt.Printf("i:%d MustReadHeaderByIncomingContext err:%+v \n", i, err)
			fmt.Println(i, " MustReadHeaderByIncomingContext header:", header)
			fmt.Printf("i:%d MustReadHeaderByIncomingContext header:%+v \n", i, header)
		})
	}
}
