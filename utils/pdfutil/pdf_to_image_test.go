package pdfutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/14 10:19
 * @Desc:
 */

func TestPdfToImageV1(t *testing.T) {
	type args struct {
		pdfUrl string
	}
	tests := []struct {
		name          string
		args          args
		wantImagePath string
		wantErr       bool
	}{
		{
			name: "",
			args: args{
				pdfUrl: "https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf",
			},
			wantImagePath: "",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotImagePath, err := PdfToImageV1(tt.args.pdfUrl)
			fmt.Println(gotImagePath, err)
		})
	}
}

func TestPdfToImageV2(t *testing.T) {
	type args struct {
		pdfUrl string
	}
	tests := []struct {
		name          string
		args          args
		wantImagePath string
		wantErr       bool
	}{
		{
			name: "",
			args: args{
				pdfUrl: "https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf",
			},
			wantImagePath: "",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotImagePath, err := PdfToImageV2(tt.args.pdfUrl)
			fmt.Println(gotImagePath, err)
		})
	}
}
