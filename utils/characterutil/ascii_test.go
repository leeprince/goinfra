package characterutil

import (
	"encoding/hex"
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/20 00:49
 * @Desc:
 */

func TestASCIICharToASCIIIntAndASCIIHex(t *testing.T) {
	type args struct {
		char int8
	}
	tests := []struct {
		name       string
		args       args
		wantIntStr string
		wantHexStr string
	}{
		{
			name: "",
			args: args{
				char: 'A',
			},
			wantIntStr: "",
			wantHexStr: "",
		},
		{
			name: "",
			args: args{
				char: 'B',
			},
			wantIntStr: "",
			wantHexStr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIntStr, gotHexStr := ASCIICharToASCIIIntAndASCIIHex(tt.args.char)
			fmt.Println("gotIntStr:", gotIntStr)
			fmt.Println("gotHexStr:", gotHexStr)
		})
	}
}

func TestASCIICharStrToASCIIIntAndASCIIHex(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name       string
		args       args
		wantIntStr string
		wantHexStr string
	}{
		{
			name: "",
			args: args{
				char: "A",
			},
			wantIntStr: "",
			wantHexStr: "",
		},
		{
			name: "",
			args: args{
				char: "B",
			},
			wantIntStr: "",
			wantHexStr: "",
		},
		{
			name: "",
			args: args{
				char: "AB",
			},
			wantIntStr: "",
			wantHexStr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIntStr, gotHexStr, err := ASCIICharStrToASCIIIntAndASCIIHex(tt.args.char)
			fmt.Println("err:", err)
			fmt.Println("gotIntStr:", gotIntStr)
			fmt.Println("gotHexStr:", gotHexStr)
		})
	}
}

func TestASCIICharStrToASCIIIntAndASCIIHexContainNoASCII(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name       string
		args       args
		wantIntStr string
		wantHexStr string
	}{
		{
			name: "",
			args: args{
				char: "A",
			},
			wantIntStr: "",
			wantHexStr: "",
		},
		{
			name: "",
			args: args{
				char: "B",
			},
			wantIntStr: "",
			wantHexStr: "",
		},
		{
			name: "",
			args: args{
				char: "AB",
			},
			wantIntStr: "",
			wantHexStr: "",
		},
		{
			name: "",
			args: args{
				char: "一等",
			},
			wantIntStr: "",
			wantHexStr: "",
		},
		{
			name: "",
			args: args{
				char: "二等",
			},
			wantIntStr: "",
			wantHexStr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(">>>>", tt.args.char)
			
			gotIntStr, gotHexStr, err := ASCIICharStrToASCIIIntAndASCIIHex(tt.args.char)
			fmt.Println("err:", err)
			fmt.Println("gotIntStr:", gotIntStr)
			fmt.Println("gotHexStr:", gotHexStr)
			
			got, err := ASCIIHexStrToASCIIChar(gotHexStr)
			fmt.Println(err)
			fmt.Println(got)
		})
	}
}

func TestIntToChar(t *testing.T) {
	type args struct {
		i int8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				i: 65,
			},
			want: "",
		},
		{
			name: "",
			args: args{
				i: 66,
			},
			want: "",
		},
		{
			name: "",
			args: args{
				i: 67,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ASCIIIntToASCIIChar(tt.args.i)
			fmt.Println(got)
		})
	}
}

func TestASCIIIntsToASCIIChars(t *testing.T) {
	type args struct {
		ii []int8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				ii: []int8{65, 66, 67},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ASCIIIntArrToASCIIChars(tt.args.ii)
			fmt.Println(got)
		})
	}
}

func TestASCIIIntStrToASCIIChars(t *testing.T) {
	type args struct {
		ii string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ii: "656667",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ASCIIIntStrToASCIIChars(tt.args.ii)
			fmt.Println(err)
			fmt.Println(got)
		})
	}
}

func TestASCIIHexToASCIIChar(t *testing.T) {
	type args struct {
		hexStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				hexStr: "4142",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ASCIIHexStrToASCIIChar(tt.args.hexStr)
			fmt.Println(err)
			fmt.Println(got)
		})
	}
}

func TestHex(t *testing.T) {
	str := "一等"
	utf8Bytes := []byte(str)
	hexStr := hex.EncodeToString(utf8Bytes)
	fmt.Println(hexStr)
	
	decodedBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println("解码失败：", err)
		return
	}
	decodedStr := string(decodedBytes)
	fmt.Println(decodedStr)
	fmt.Println("---------")
}
