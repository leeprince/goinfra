package characterutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/22 15:26
 * @Desc:
 */

func TestStringToHex(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "A",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "B",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "AB",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "一等",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "二等",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hexStr := StringToHexStr(tt.args.s)
			fmt.Println("StringToHexStr:", hexStr)
			
			str, err := HexStrToString(hexStr)
			fmt.Println("HexStrToString err:", err)
			fmt.Println("HexStrToString str:", str)
			
			fmt.Println("---")
			gotIntStr, gotHexStr, err := ASCIICharStrToASCIIIntAndASCIIHex(tt.args.s)
			fmt.Println("ASCIICharStrToASCIIIntAndASCIIHex err:", err)
			fmt.Println("ASCIICharStrToASCIIIntAndASCIIHex gotIntStr:", gotIntStr)
			fmt.Println("ASCIICharStrToASCIIIntAndASCIIHex gotHexStr:", gotHexStr)
			
			got, err := ASCIIHexStrToASCIIChar(gotHexStr)
			fmt.Println("ASCIIHexStrToASCIIChar", err)
			fmt.Println("ASCIIHexStrToASCIIChar", got)
		})
	}
}
