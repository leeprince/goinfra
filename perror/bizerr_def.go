package code

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/23 下午11:23
 * @Desc:   通用状态码及描述
 */

var (
	// 特殊的状态码
	BizErrSuccess           = NewBizErr(0, "成功")
	BizErrNotAuthentication = NewBizErr(401, "未授权!")
	BizErrAccessForbidden   = NewBizErr(403, "禁止访问!")

	BizErrRequired      = NewBizErr(100000001, "缺少必填项!")
	BizErrFormat        = NewBizErr(100000002, "格式错误!")
	BizErrConvertFormat = NewBizErr(100000003, "格式转换错误!")
	BizErrLen           = NewBizErr(100000004, "长度错误!")
	BizErrDataEmpty     = NewBizErr(100000005, "数据为空!")
	BizErrTimeOut       = NewBizErr(100000006, "请求超时!")
	BizErrVerify        = NewBizErr(100000015, "验证失败!")
	BizErrTypeAsserts   = NewBizErr(100000007, "类型断言失败!")
	BizErrNoExistType   = NewBizErr(100000008, "类型不存在!")
	BizErrNoExistSource = NewBizErr(100000009, "来源不存在!")
	BizErrNoExistMethod = NewBizErr(100000010, "方法不存在!")
	BizErrNoExistFile   = NewBizErr(100000011, "文件不存在!")
	BizErrEncrypt       = NewBizErr(100000012, "加密失败!")
	BizErrDecrypt       = NewBizErr(100000013, "解密失败!")
	BizErrSign          = NewBizErr(100000014, "签名失败!")
	BizErrVerifySign    = NewBizErr(100000015, "签名验证失败!")
	BizErrGenerateData  = NewBizErr(100000016, "生成数据错误!")
	BizErrEncode        = NewBizErr(100000017, "编码错误!")
	BizErrDecode        = NewBizErr(100000018, "解码错误!")
	BizErrPanic         = NewBizErr(100000999, "处理异常!")
)
