package wechatopen

import (
	"fmt"
	"github.com/leeprince/goinfra/http/httpcli"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/6/16 11:29
 * @Desc:
 */

type UserInfo struct {
	Openid     string `json:"openid"`     // 	普通用户的标识，对当前开发者账号唯一
	Nickname   string `json:"nickname"`   // 	普通用户昵称
	Sex        string `json:"sex"`        // 	普通用户性别，1为男性，2为女性
	Province   string `json:"province"`   // 	普通用户个人资料填写的省份
	City       string `json:"city"`       // 	普通用户个人资料填写的城市
	Country    string `json:"country"`    // 	国家，如中国为CN
	Headimgurl string `json:"headimgurl"` // 	用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空
	Privilege  string `json:"privilege"`  // 	用户特权信息，json数组，如微信沃卡用户为（chinaunicom）
	Unionid    string `json:"unionid"`    // 	用户统一标识。针对一个微信开放平台账号下的应用，同一用户的unionid是唯一的。
}

// GetUserInfo 获取用户个人信息（UnionID机制）
// 官方文档：https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Authorized_Interface_Calling_UnionID.html
func (c *WechatOpenSDK) GetUserInfo(accessToken, openid string) (resp *UserInfo, err error) {
	resp = &UserInfo{}
	
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
		accessToken, openid)
	cli := httpcli.NewHttpClient()
	_, _, err = cli.WithNotLogging(false).
		WithMethod(http.MethodGet).
		WithURL(url).
		WithResponse(resp).
		Do()
	
	return
}
