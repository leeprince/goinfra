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
    BizErrLen           = NewBizErr(1000003, "长度错误!")
    BizErrDataEmpty     = NewBizErr(1000004, "数据为空!")
    BizErrTimeOut       = NewBizErr(1000005, "请求超时!")
    BizErrNoExistType   = NewBizErr(1000006, "类型不存在!")
    BizErrNoExistSource = NewBizErr(1000007, "来源不存在!")
    BizErrNoExistMethod = NewBizErr(1000008, "方法不存在!")
    BizErrNoExistFile   = NewBizErr(1000008, "文件不存在!")
)
