package perror

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/23 下午11:23
 * @Desc:   通用状态码及描述
				状态码格式： 微服务(3位)+业务逻辑标识(3位)
    				- 特殊的状态码：正确响应的状态码等于0(非0都表示为处理失败)、未授权、禁止访问
					- 该通用状态码的格式为：微服务(通用服务:100)+业务逻辑标识(xxx)
*/

var (
	// 特殊的状态码
	Success = NewBizErr(0, "成功")
	Fail    = NewBizErr(1, "系统异常，请重试！")
	
	BizErrNotAuthentication = NewBizErr(401, "未授权!")
	BizErrAccessForbidden   = NewBizErr(403, "禁止访问!")
	
	BizErrRequired = NewBizErr(100001, "缺少必填项!")
	BizErrLen      = NewBizErr(100002, "长度错误!")
	
	BizErrFormat        = NewBizErr(100003, "格式错误!")
	BizErrFormatConvert = NewBizErr(100004, "格式转换错误!")
	
	BizErrDataEmpty    = NewBizErr(100005, "数据为空!")
	BizErrDataNil      = NewBizErr(100006, "数据为nil!")
	BizErrDataExist    = NewBizErr(100007, "数据已存在!")
	BizErrDataParse    = NewBizErr(100008, "数据解析错误!")
	BizErrDataGenerate = NewBizErr(100009, "数据生成错误!")
	
	BizErrRequestTimeOut  = NewBizErr(100010, "请求超时!")
	BizErrRequestErr      = NewBizErr(100011, "请求失败!")
	BizErrRequestCodeFail = NewBizErr(100012, "请求响应code非成功!")
	
	BizErrTypeAsserts = NewBizErr(100013, "类型断言失败!")
	BizErrTypeNoExist = NewBizErr(100014, "类型不存在!")
	
	BizErrSourceNoExist = NewBizErr(100015, "来源不存在!")
	
	BizErrMethodNoExist = NewBizErr(100016, "方法不存在!")
	
	BizErrFileNoExist = NewBizErr(100017, "文件不存在!")
	
	BizErrSecurityEncrypt      = NewBizErr(100018, "加密失败!")
	BizErrSecurityDecrypt      = NewBizErr(100019, "解密失败!")
	BizErrSecurityEncode       = NewBizErr(100020, "编码错误!")
	BizErrSecurityDecode       = NewBizErr(100021, "解码错误!")
	BizErrSecuritySignGenerate = NewBizErr(100022, "生成签名失败!")
	BizErrSecuritySignVerify   = NewBizErr(100023, "验证签名失败!")
	BizErrSecurityAuthGenerate = NewBizErr(100024, "生成授权失败!")
	BizErrSecurityAuthVerify   = NewBizErr(100025, "未授权，请联系运营商!")
	BizErrSecurityAuthing      = NewBizErr(100026, "授权中，请联系运营商!")
	BizErrSecurityAuthExpired  = NewBizErr(100027, "授权已过期，请联系运营商!")
	
	BizErrOpreate    = NewBizErr(100028, "操作失败!")
	BizErrInsert     = NewBizErr(100029, "插入失败!")
	BizErrDelete     = NewBizErr(100030, "删除失败!")
	BizErrUpdate     = NewBizErr(100031, "更新失败!")
	BizErrFind       = NewBizErr(100032, "查询失败!")
	BizErrInsertList = NewBizErr(100033, "插入列表失败!")
	BizErrDeleteList = NewBizErr(100034, "删除列表失败!")
	BizErrUpdateList = NewBizErr(100035, "更新列表失败!")
	BizErrFindList   = NewBizErr(100036, "查询列表失败!")
	
	BizErrPub     = NewBizErr(100037, "发布失败!")
	BizErrConsume = NewBizErr(100038, "消费失败!")
	
	BizErrPanic = NewBizErr(100099, "处理异常!")
)
