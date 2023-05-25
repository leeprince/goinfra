package fileutil

import (
	"fmt"
	"os"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/24 18:15
 * @Desc:
 */

func TestCreateFileAndWrite(t *testing.T) {
	type args struct {
		path      string
		pathFile  string
		fileBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				path:      "f:\\tmp",
				pathFile:  "f:\\tmp\\hello.txt",
				fileBytes: []byte("hello"),
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				path:      fmt.Sprintf("f:%s%s", string(os.PathSeparator), "tmp"),
				pathFile:  fmt.Sprintf("f:%s%s%s%s", string(os.PathSeparator), "tmp", string(os.PathSeparator), "hello.txt"),
				fileBytes: []byte("hello world"),
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				path:      fmt.Sprintf("f:%s%s%s", string(os.PathSeparator), "tmp", string(os.PathSeparator)),
				pathFile:  fmt.Sprintf("f:%s%s%s%s", string(os.PathSeparator), "tmp", string(os.PathSeparator), "hello.txt"),
				fileBytes: []byte("hello world"),
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				path:      "F:\\tmp\\e-invoice-invoice-ctl\\local\\",
				pathFile:  "F:\\tmp\\e-invoice-invoice-ctl\\local\\144e5129b66743c2_7025391305169615131.pdf",
				fileBytes: []byte("hello world"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFileAndWrite(tt.args.path, tt.args.pathFile, tt.args.fileBytes); (err != nil) != tt.wantErr {
				t.Errorf("CreateFileAndWrite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
