package code

import "fmt"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/23 下午10:31
 * @Desc:   BizErr 错误
 */

type BizErr struct {
    code    int32
    message string
}

func NewBizErr(code int32, message string) BizErr {
    return BizErr{
        code:    code,
        message: message,
    }
}

func (e BizErr) Error() string {
    return fmt.Sprintf("%d:%s", e.code, e.message)
}

func (e BizErr) GetCode() int32 {
    return e.code
}

func (e BizErr) GetMessage() string {
    return e.message
}

