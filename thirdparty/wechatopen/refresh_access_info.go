package wechatopen

import (
	"fmt"
	"github.com/leeprince/goinfra/http/httpcli"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/6/16 11:22
 * @Desc:
 */

// RefreshAccessToken 刷新或续期access_token使用
// 官方文档：https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Authorized_Interface_Calling_UnionID.html
func (c *WechatOpenSDK) RefreshAccessToken(refreshToken string) (resp *AccessTokenInfo, err error) {
	resp = &AccessTokenInfo{}
	
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
		c.appid, c.secret, refreshToken)
	cli := httpcli.NewHttpClient()
	_, _, err = cli.WithNotLogging(false).
		WithMethod(http.MethodGet).
		WithURL(url).
		WithResponse(resp).
		Do()
	
	return
}
