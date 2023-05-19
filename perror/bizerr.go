package perror

import (
	"fmt"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/23 下午10:31
 * @Desc:   BizErr 错误
 */

type BizErr struct {
	code    int32
	message string
	error   error
}

func NewBizErr(code int32, message string) BizErr {
	return BizErr{
		code:    code,
		message: message,
	}
}

func (e BizErr) Error() string {
	if e.error == nil {
		return fmt.Sprintf("%d:%s", e.code, e.message)
	}
	return e.error.Error()
}

// 添加错误信息。
//   - 通过`errors.Unwrap(err)`解错误包裹层得到最后一次添加的error，中间的error包裹每次都通过`e.Error()`屏蔽了。
//   - msgs: 只取msgs[0]
func (e BizErr) WithField(msg string) BizErr {
	e.message = e.message + "#" + msg
	return e
}

// 添加错误信息。
//   - 通过`errors.Unwrap(err)`解错误包裹层得到最后一次添加的error，中间的error包裹每次都通过`e.Error()`屏蔽了。
//   - msgs: 只取msgs[0]
func (e BizErr) WithError(err error, msgs ...string) BizErr {
	if err == nil {
		return e
	}
	if len(msgs) > 0 && msgs[0] != "" {
		e.error = fmt.Errorf(e.Error()+"(%s:%w)", msgs[0], err)
		return e
	}
	e.error = fmt.Errorf(e.Error()+"(%w)", err)
	return e
}

func (e BizErr) GetCode() int32 {
	return e.code
}

func (e BizErr) GetMessage() string {
	return e.message
}

// 获取错误信息。只有通过`WithError`方法添加错误信息才会返回error信息
func (e BizErr) GetError() error {
	return e.error
}
