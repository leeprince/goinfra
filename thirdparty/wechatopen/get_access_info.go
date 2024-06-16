package wechatopen

import (
	"fmt"
	"github.com/leeprince/goinfra/http/httpcli"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/6/16 03:32
 * @Desc:
 */

type AccessTokenInfo struct {
	AccessToken  string `json:"access_token"`  //  	接口调用凭证
	ExpiresIn    int32  `json:"expires_in"`    //  	access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` //  	用户刷新access_token
	Openid       string `json:"openid"`        //  	授权用户唯一标识
	Scope        string `json:"scope"`         //  	用户授权的作用域，使用逗号（,）分隔
	Unionid      string `json:"unionid"`       //  	用户统一标识。针对一个微信开放平台账号下的应用，同一用户的 unionid 是唯一的
}

// GetAccessTokenInfo 通过code获取access_token
// 官方文档：https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Authorized_Interface_Calling_UnionID.html
func (c *WechatOpenSDK) GetAccessTokenInfo(code string) (resp *AccessTokenInfo, err error) {
	resp = &AccessTokenInfo{}
	
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		c.appid, c.secret, code)
	cli := httpcli.NewHttpClient()
	_, _, err = cli.WithNotLogging(false).
		WithMethod(http.MethodGet).
		WithURL(url).
		WithResponse(resp).
		Do()
	
	return
}
