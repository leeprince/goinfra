package main

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/8 02:20
 * @Desc:
 */

func Test_getFileNameFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				path: "文章/typora/错过阿里，冲腾讯吧/640-20240306170246442.png",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				path: "/文章/typora/xxx/0001.png",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFileNameFromPath(tt.args.path)
			fmt.Println("getFileNameFromPath:", got)
			
			remoteFileDir, remoteFilePath := getFilePartPathFromPath(tt.args.path, getFilePartPathFromPathDelimiter)
			fmt.Println("remoteFileDir:", remoteFileDir)
			fmt.Println("remoteFilePath:", remoteFilePath)
		})
	}
}
