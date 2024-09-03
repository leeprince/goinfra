package wechatopen

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/6/16 02:36
 * @Desc:	微信开放平台SDK
 * 			官方文档：https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
 */

type WechatOpenSDK struct {
	appid  string
	secret string
}

func NewWechatOpenSDK(appid string, secret string) *WechatOpenSDK {
	return &WechatOpenSDK{
		appid:  appid,
		secret: secret,
	}
}
