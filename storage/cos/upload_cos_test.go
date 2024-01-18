package cos

import (
	"fmt"
	"github.com/leeprince/goinfra/config"
	"github.com/leeprince/goinfra/utils/fileutil"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"os"
	"testing"
	"time"
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
			/*if err := CreateFileAndWrite(tt.args.path, tt.args.pathFile, tt.args.fileBytes); (err != nil) != tt.wantErr {
				t.Errorf("CreateFileAndWrite() error = %v, wantErr %v", err, tt.wantErr)
			}*/

			ok, err := fileutil.WriteFile(tt.args.path, "helloword.txt", tt.args.fileBytes, true)
			fmt.Println(ok, err)
		})
	}
}

func TestUploadCos(t *testing.T) {
	config.InitConfig("../../config/config.yaml")

	cosConfig := config.C.COS
	cosClient := NewCosClient(cosConfig.SecretID, cosConfig.SecretKey, cosConfig.AppID, cosConfig.Bucket, cosConfig.Region)

	//fileContentBytes := []byte("prince")
	//cosName := "princetest" // 文本

	cosName := "test/princetest.jpeg" // 图片
	fileContentBytes, err := fileutil.ReadFile(".", "test.png")
	if err != nil {
		log.Fatal("fileutil.ReadFile:", err)
	}

	customeAccessHost := ""

	opts := &cos.ObjectPutOptions{
		ACLHeaderOptions: nil,
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentDisposition: fmt.Sprintf("attachment; filename=\"%s\"", cosName),
		},
	}
	cosUrl, err := cosClient.UploadCos(fileContentBytes, cosName, customeAccessHost, opts)
	//cosUrl, err := cosClient.UploadCos(fileContentBytes, cosName, customeAccessHost, nil)
	fmt.Println("err:", err)

	fmt.Println("cosUrl:", cosUrl)

}
func TestGetPresignedURL(t *testing.T) {
	config.InitConfig("../../config/config.yaml")

	cosConfig := config.C.COS
	cosClient := NewCosClient(cosConfig.SecretID, cosConfig.SecretKey, cosConfig.AppID, cosConfig.Bucket, cosConfig.Region)

	cosName := "test/princetest.jpeg"
	exipired := time.Second * 10
	presigndUrl, err := cosClient.GetPresignedURL(cosName, exipired)
	fmt.Println("err:", err)

	fmt.Println("presigndUrl:", presigndUrl)

}
