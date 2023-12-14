package testdata

import (
	"errors"
	"fmt"
	"github.com/leeprince/goinfra/consts/constval"
	"github.com/leeprince/goinfra/perror"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/20 02:06
 * @Desc:
 */

var (
	// 处理成功:0
	Success = constval.NewInt32(0, "成功响应")

	// 基础错误:11xxxx
	BaseErrParamsInvalid    = constval.NewInt32(110001, "无效参数!")
	BaseErrParamsRequired   = constval.NewInt32(110002, "参数必填!")
	BaseErrDataNull         = constval.NewInt32(110003, "数据为空!")
	BaseErrFind             = constval.NewInt32(110004, "查询错误!")
	BaseErrSave             = constval.NewInt32(110005, "保存错误!")
	BaseErrDelete           = constval.NewInt32(110006, "删除错误!")
	BaseErrSign             = constval.NewInt32(110007, "签名错误!")
	BaseErrGetAccessToken   = constval.NewInt32(110008, "获取access_token失败!")
	BaseErrCheckAccessToken = constval.NewInt32(110009, "检查access_token失败!")
	BaseErrConfig           = constval.NewInt32(110010, "配置错误!")
	BaseErrDataParse        = constval.NewInt32(110011, "数据解析错误!")
	BaseErrEventPublish     = constval.NewInt32(110012, "事件发布失败!")
	BaseErrLimit            = constval.NewInt32(110013, "检查限流错误!")
	BaseErrFieldType        = constval.NewInt32(110014, "字段类型转换错误!")
	BaseErrInstance         = constval.NewInt32(110015, "获取实例失败!")
	BaseErrGetToken         = constval.NewInt32(110016, "获取token失败!")
	BaseErrCheckToken       = constval.NewInt32(110017, "检查token失败!")
	BaseErrGetPassword      = constval.NewInt32(110018, "获取密码错误!")
	BaseErrSetPassword      = constval.NewInt32(110019, "设置密码错误!")
)

func TestReturnError(t *testing.T) {
	var err error
	err = ReturnError()
	fmt.Println(err)
}
func ReturnError() error {
	return perror.NewBizErr(Success.Key(), Success.Value())
}

func TestReturnErrorError(t *testing.T) {
	var err error
	err = ReturnErrorError()
	fmt.Println(err)
}
func ReturnErrorError() error {
	err := errors.New(">ReturnErrorError 01")
	return perror.NewBizErr(Success.Key(), Success.Value()).WithError(err)
}

func TestReturnErrorErrorMsg(t *testing.T) {
	var err error
	err = ReturnErrorErrorMsg()
	fmt.Println(err)
}
func ReturnErrorErrorMsg() error {
	err := errors.New(">ReturnErrorError 01")
	return perror.NewBizErr(Success.Key(), Success.Value()).WithError(err, "Msg")
}

func TestReturnErrorErrorMsgNesting(t *testing.T) {
	var err error
	err = ReturnErrorErrorMsgNesting1()
	fmt.Println(err)

	err = perror.NewBizErr(BaseErrParamsInvalid.Key(), BaseErrParamsInvalid.Value()).WithError(err, "嵌套1")
	fmt.Println(err)

	err = perror.NewBizErr(BaseErrParamsRequired.Key(), BaseErrParamsRequired.Value()).WithError(err, "嵌套2")
	fmt.Println(err)
}
func ReturnErrorErrorMsgNesting1() error {
	err := errors.New(">ReturnErrorErrorMsgNesting1")
	return perror.NewBizErr(Success.Key(), Success.Value()).WithError(err, "Msg")
}
