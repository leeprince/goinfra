package botutil

import (
	"net/http"
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/16 20:16
 * @Desc:
 */

const botUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxxxx"

func TestSendQYWXBot(t *testing.T) {
	type args struct {
		url         string
		contextType BotContentType
		title       string
		contents    []string
	}
	tests := []struct {
		name              string
		args              args
		wantRespBobyBytes []byte
		wantResp          *http.Response
		wantErr           bool
	}{
		{
			name: "",
			args: args{
				url:         botUrl,
				contextType: BOT_CONTENTTYPE_TEXT,
				title:       "登录次数超过限制",
				contents: []string{
					"税号：91441302MA53J2H49U",
				},
			},
			wantRespBobyBytes: nil,
			wantResp:          nil,
			wantErr:           false,
		},
		{
			name: "",
			args: args{
				url:         botUrl,
				contextType: BOT_CONTENTTYPE_MARKDOWN,
				title:       "登录次数超过限制",
				contents: []string{
					"税号：91441302MA53J2H49U",
				},
			},
			wantRespBobyBytes: nil,
			wantResp:          nil,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRespBobyBytes, gotResp, err := SendQYWXBot(tt.args.url, tt.args.contextType, tt.args.title, tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendQYWXBot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespBobyBytes, tt.wantRespBobyBytes) {
				t.Errorf("SendQYWXBot() gotRespBobyBytes = %v, want %v", gotRespBobyBytes, tt.wantRespBobyBytes)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SendQYWXBot() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
