package perror

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/23 下午11:23
 * @Desc:   通用状态码及描述
				状态码格式： 业务线(2位)+微服务(2位)+功能模块(2位)+业务逻辑标识(3位)
    				- 特殊的状态码：正确响应的状态码等于0(非0都表示为处理失败)、未授权、禁止访问
					- 该通用状态码的格式为：业务线(业务线通用10)+微服务(通用服务:00)+功能模块(通用功能：00)+业务逻辑标识(xxx)
*/

var (
	// 特殊的状态码
	Success = NewBizErr(0, "成功")
	Fail    = NewBizErr(1, "系统异常，请重试！")

	BizErrNotAuthentication = NewBizErr(401, "未授权!")
	BizErrAccessForbidden   = NewBizErr(403, "禁止访问!")

	BizErrRequired = NewBizErr(100001001, "缺少必填项!")
	BizErrLen      = NewBizErr(100001002, "长度错误!")

	BizErrFormat        = NewBizErr(100002001, "格式错误!")
	BizErrFormatConvert = NewBizErr(100002002, "格式转换错误!")

	BizErrDataEmpty    = NewBizErr(100003001, "数据为空!")
	BizErrDataNil      = NewBizErr(100003002, "数据为nil!")
	BizErrDataExist    = NewBizErr(100003003, "数据已存在!")
	BizErrDataParse    = NewBizErr(100003004, "数据解析错误!")
	BizErrDataGenerate = NewBizErr(100003005, "数据生成错误!")

	BizErrRequestTimeOut  = NewBizErr(100004001, "请求超时!")
	BizErrRequestErr      = NewBizErr(100004002, "请求失败!")
	BizErrRequestCodeFail = NewBizErr(100004003, "请求响应code非成功!")

	BizErrTypeAsserts = NewBizErr(100005001, "类型断言失败!")
	BizErrTypeNoExist = NewBizErr(100005002, "类型不存在!")

	BizErrSourceNoExist = NewBizErr(100006001, "来源不存在!")

	BizErrMethodNoExist = NewBizErr(100007001, "方法不存在!")

	BizErrFileNoExist = NewBizErr(100008001, "文件不存在!")

	BizErrSecurityEncrypt      = NewBizErr(100009001, "加密失败!")
	BizErrSecurityDecrypt      = NewBizErr(100009002, "解密失败!")
	BizErrSecurityEncode       = NewBizErr(100009003, "编码错误!")
	BizErrSecurityDecode       = NewBizErr(100009004, "解码错误!")
	BizErrSecuritySignGenerate = NewBizErr(100009005, "生成签名失败!")
	BizErrSecuritySignVerify   = NewBizErr(100009006, "验证签名失败!")
	BizErrSecurityAuthGenerate = NewBizErr(100009006, "生成授权失败!")
	BizErrSecurityAuthVerify   = NewBizErr(100009006, "未授权，请联系运营商!")
	BizErrSecurityAuthing      = NewBizErr(100009006, "授权中，请联系运营商!")
	BizErrSecurityAuthExpired  = NewBizErr(100009006, "授权已过期，请联系运营商!")

	BizErrOpreate    = NewBizErr(100010001, "操作失败!")
	BizErrInsert     = NewBizErr(100010002, "插入失败!")
	BizErrDelete     = NewBizErr(100010003, "删除失败!")
	BizErrUpdate     = NewBizErr(100010004, "更新失败!")
	BizErrFind       = NewBizErr(100010005, "查询失败!")
	BizErrInsertList = NewBizErr(100010006, "插入列表失败!")
	BizErrDeleteList = NewBizErr(100010007, "删除列表失败!")
	BizErrUpdateList = NewBizErr(100010008, "更新列表失败!")
	BizErrFindList   = NewBizErr(100010009, "查询列表失败!")

	BizErrPub     = NewBizErr(100011001, "发布失败!")
	BizErrConsume = NewBizErr(100011002, "消费失败!")

	BizErrPanic = NewBizErr(101099999, "处理异常!")
)
