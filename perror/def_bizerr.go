package perror

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/23 下午11:23
 * @Desc:   通用状态码及描述
 */

var (
	// 特殊的状态码
	Success = NewBizErr(0, "成功")
	Fail    = NewBizErr(1, "系统异常，请重试！")
	
	BizErrNotAuthentication = NewBizErr(401, "未授权!")
	BizErrAccessForbidden   = NewBizErr(403, "禁止访问!")
	
	BizErrRequired             = NewBizErr(100001001, "缺少必填项!")
	BizErrLen                  = NewBizErr(100001004, "长度错误!")
	BizErrFormat               = NewBizErr(100002002, "格式错误!")
	BizErrFormatConvert        = NewBizErr(100002003, "格式转换错误!")
	BizErrDataEmpty            = NewBizErr(100003005, "数据为空!")
	BizErrDataNil              = NewBizErr(100003006, "数据为nil!")
	BizErrDataExist            = NewBizErr(100003007, "数据已存在!")
	BizErrDataParse            = NewBizErr(100003008, "数据解析错误!")
	BizErrDataGenerate         = NewBizErr(100003009, "数据生成错误!")
	BizErrRequestTimeOut       = NewBizErr(100004010, "请求超时!")
	BizErrRequestErr           = NewBizErr(100004011, "请求失败!")
	BizErrRequestCodeFail      = NewBizErr(100004011, "请求响应code非成功!")
	BizErrTypeAsserts          = NewBizErr(100005012, "类型断言失败!")
	BizErrTypeNoExist          = NewBizErr(100005013, "类型不存在!")
	BizErrSourceNoExist        = NewBizErr(100006014, "来源不存在!")
	BizErrMethodNoExist        = NewBizErr(100007015, "方法不存在!")
	BizErrFileNoExist          = NewBizErr(100008016, "文件不存在!")
	BizErrSecurityEncrypt      = NewBizErr(100009017, "加密失败!")
	BizErrSecurityDecrypt      = NewBizErr(100009018, "解密失败!")
	BizErrSecurityEncode       = NewBizErr(100009021, "编码错误!")
	BizErrSecurityDecode       = NewBizErr(100009022, "解码错误!")
	BizErrSecuritySignGenerate = NewBizErr(100009019, "签名生成失败!")
	BizErrSignVerify           = NewBizErr(100009020, "签名验证失败!")
	BizErrOpreate              = NewBizErr(100010024, "操作失败!")
	BizErrPanic                = NewBizErr(101099999, "处理异常!")
)
