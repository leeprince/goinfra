package wechatopen

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/6/16 02:36
 * @Desc:
 */

type WechatClient struct {
	appid  string
	secret string
}

func NewWechatClient(appid string, secret string) *WechatClient {
	return &WechatClient{
		appid:  appid,
		secret: secret,
	}
}
