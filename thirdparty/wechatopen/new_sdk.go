package wechatopen

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/6/16 02:36
 * @Desc:
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
