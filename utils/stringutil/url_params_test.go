package stringutil

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/13 14:13
 * @Desc:
 */
func TestGetUrlStrParams(t *testing.T) {
	type args struct {
		urlString string
		key       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				urlString: "https://localhost?optype=BookingOrderForAli&ExistOrderNumbers=1",
				key:       "optype",
			},
			want: "BookingOrderForAli",
		},
		{
			name: "",
			args: args{
				// \u0026 => &
				urlString: "http://localhost?optype=BookingOrderForAli\u0026ExistOrderNumbers=1",
				key:       "optype",
			},
			want: "BookingOrderForAli",
		},
		{
			name: "",
			args: args{
				urlString: "optype=BookingOrderForAli\u0026ExistOrderNumbers=1",
				key:       "optype",
			},
			want: "BookingOrderForAli",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := GetUrlStrParam(tt.args.urlString, tt.args.key); got != tt.want || err != nil {
				t.Errorf("GetUrlStrParam() = %v, want %v, err %v", got, tt.want, err)
			} else {
				fmt.Println(got)
			}
		})
	}
}

func TestGetUrlStrParams1(t *testing.T) {
	type args struct {
		urlString string
		keys      []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				// \u0026 => &
				urlString: "http://localhost?optype=BookingOrderForAli&ExistOrderNumbers=1",
				keys:      []string{"optype", "ExistOrderNumbers"},
			},
			want: map[string]string{
				"optype":            "BookingOrderForAli",
				"ExistOrderNumbers": "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUrlStrParams(tt.args.urlString, tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUrlStrParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUrlStrParams() got = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
