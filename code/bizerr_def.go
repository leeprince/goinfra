package code

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/23 下午11:23
 * @Desc:   通用状态码及描述
 *              - 状态码格式： 微服务(2位)+功能模块(2位)+业务逻辑标识(3位)
 *                  - 特殊的状态码：正确响应的状态码等于0
 *                  - 该通用状态码的格式为：微服务(通用服务:10)+功能模块(通用功能：00)+业务逻辑标识(xxx)
 */

var (
    // 特殊的状态码：正确响应的状态码等于0
    BizErrSuccess = NewBizErr(0, "success")
    
    BizErrRequired      = NewBizErr(1000001, "缺少必填项!")
    BizErrFormat        = NewBizErr(1000002, "格式错误!")
    BizErrConvertFormat = NewBizErr(1000003, "格式转换错误!")
    BizErrLen           = NewBizErr(1000004, "长度错误!")
    BizErrDataEmpty     = NewBizErr(1000005, "数据为空!")
    BizErrTimeOut       = NewBizErr(1000006, "请求超时!")
    BizErrVerify        = NewBizErr(1000015, "验证失败!")
    BizErrTypeAsserts   = NewBizErr(1000007, "类型断言失败!")
    BizErrNoExistType   = NewBizErr(1000008, "类型不存在!")
    BizErrNoExistSource = NewBizErr(1000009, "来源不存在!")
    BizErrNoExistMethod = NewBizErr(1000010, "方法不存在!")
    BizErrNoExistFile   = NewBizErr(1000011, "文件不存在!")
    BizErrEncrypt       = NewBizErr(1000012, "加密失败!")
    BizErrDecrypt       = NewBizErr(1000013, "解密失败!")
    BizErrSign          = NewBizErr(1000014, "签名失败!")
    BizErrVerifySign    = NewBizErr(1000015, "签名验证失败!")
    BizErrGenerateData  = NewBizErr(1000016, "生成数据错误!")
)
