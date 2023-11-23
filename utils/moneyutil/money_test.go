package moneyutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/13 23:02
 * @Desc:
 */

func TestJiaoToYuan(t *testing.T) {
	type args struct {
		jiao int64
	}
	tests := []struct {
		name     string
		args     args
		wantYuan string
	}{
		{
			name: "",
			args: args{
				jiao: 1234,
			},
			wantYuan: "123.4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotYuan := JiaoToYuan(tt.args.jiao); gotYuan != tt.wantYuan {
				t.Errorf("JiaoToYuan() = %v, want %v", gotYuan, tt.wantYuan)
			}
		})
	}
}

func TestYuanToJiao(t *testing.T) {
	type args struct {
		yuan int64
	}
	tests := []struct {
		name     string
		args     args
		wantJiao int64
	}{
		{
			name: "",
			args: args{
				yuan: 1234,
			},
			wantJiao: 12340,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJiao := YuanToJiao(tt.args.yuan); gotJiao != tt.wantJiao {
				t.Errorf("YuanToJiao() = %v, want %v", gotJiao, tt.wantJiao)
			}
		})
	}
}

func TestFenToYuan(t *testing.T) {
	type args struct {
		fen int64
	}
	tests := []struct {
		name     string
		args     args
		wantYuan string
	}{
		{
			name: "",
			args: args{
				fen: 10000,
			},
			wantYuan: "100",
		},
		{
			name: "",
			args: args{
				fen: 10010,
			},
			wantYuan: "100.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotYuan := FenToYuan(tt.args.fen); gotYuan != tt.wantYuan {
				t.Errorf("FenToYuan() = %v, want %v", gotYuan, tt.wantYuan)
			}
		})
	}
}

func TestYuanToFen(t *testing.T) {
	type args struct {
		yuan string
	}
	tests := []struct {
		name    string
		args    args
		wantFen int64
	}{

		{
			name: "",
			args: args{
				yuan: "101",
			},
			wantFen: 10100,
		},

		{
			name: "",
			args: args{
				yuan: "10.1",
			},
			wantFen: 10100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fen, err := YuanToFen(tt.args.yuan)
			fmt.Printf("fen:%+v---err:%+v \n", fen, err)
		})
	}
}

func TestFenFloat64ToYuan(t *testing.T) {
	type args struct {
		fen float64
	}
	tests := []struct {
		name     string
		args     args
		wantYuan string
	}{
		{
			name: "",
			args: args{
				fen: 10000,
			},
			wantYuan: "100",
		},
		{
			name: "",
			args: args{
				fen: 12345,
			},
			wantYuan: "123.45",
		},
		{
			name: "",
			args: args{
				fen: 12345678,
			},
			wantYuan: "123456.78",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotYuan := FenFloat64ToYuan(tt.args.fen); gotYuan != tt.wantYuan {
				t.Errorf("FenFloat64ToYuan() = %v, want %v", gotYuan, tt.wantYuan)
			}
		})
	}
}

func TestFenFloat64ToCeilYuan(t *testing.T) {
	type args struct {
		fen float64
	}
	tests := []struct {
		name     string
		args     args
		wantYuan string
	}{
		{
			name: "",
			args: args{
				fen: 10000,
			},
			wantYuan: "100",
		},
		{
			name: "",
			args: args{
				fen: 12345,
			},
			wantYuan: "124",
		},
		{
			name: "",
			args: args{
				fen: 12345678,
			},
			wantYuan: "123457",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotYuan := FenFloat64ToCeilYuan(tt.args.fen); gotYuan != tt.wantYuan {
				t.Errorf("FenFloat64ToYuan() = %v, want %v", gotYuan, tt.wantYuan)
			}
		})
	}
}
func TestFenFloat64ToFloorYuan(t *testing.T) {
	type args struct {
		fen float64
	}
	tests := []struct {
		name     string
		args     args
		wantYuan string
	}{
		{
			name: "",
			args: args{
				fen: 10000,
			},
			wantYuan: "100",
		},
		{
			name: "",
			args: args{
				fen: 12345,
			},
			wantYuan: "123",
		},
		{
			name: "",
			args: args{
				fen: 12345678,
			},
			wantYuan: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotYuan := FenFloat64ToFloorYuan(tt.args.fen); gotYuan != tt.wantYuan {
				t.Errorf("FenFloat64ToYuan() = %v, want %v", gotYuan, tt.wantYuan)
			}
		})
	}
}

func TestFenFloat64ToRoundYuan(t *testing.T) {
	type args struct {
		fen float64
	}
	tests := []struct {
		name     string
		args     args
		wantYuan string
	}{
		{
			name: "",
			args: args{
				fen: 10000,
			},
			wantYuan: "100",
		},
		{
			name: "",
			args: args{
				fen: 12345,
			},
			wantYuan: "123.45",
		},
		{
			name: "",
			args: args{
				fen: 12345678,
			},
			wantYuan: "123456.78",
		},
		{
			name: "",
			args: args{
				fen: 123.45678,
			},
			wantYuan: "1.23",
		},
		{
			name: "",
			args: args{
				fen: 1234.5678,
			},
			wantYuan: "12.35",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotYuan := FenFloat64ToRoundYuan(tt.args.fen); gotYuan != tt.wantYuan {
				t.Errorf("FenFloat64ToYuan() = %v, want %v", gotYuan, tt.wantYuan)
			}
		})
	}
}
