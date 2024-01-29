package fileutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/14 10:01
 * @Desc:
 */

func TestGetFileInfoByUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantInfo FileInfoOfUrl
	}{
		{
			name: "",
			args: args{
				url: "https://example.com/images/cat.jpg",
			},
			wantInfo: FileInfoOfUrl{},
		},
		{
			name: "",
			args: args{
				url: "https://example.com/images/cat",
			},
			wantInfo: FileInfoOfUrl{},
		},
		{
			name: "",
			args: args{
				url: "https://example.com/images/cat.bak.jpg",
			},
			wantInfo: FileInfoOfUrl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo := GetFileInfoByUrl(tt.args.url)
			fmt.Println(gotInfo)
		})
	}
}

func TestGetFilePathOfName(t *testing.T) {
	fmt.Println(GetFilePathOfName("./dd.png"))
	fmt.Println(GetFilePathOfName("dd.png"))
}
