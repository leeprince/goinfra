package kv

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/26 下午1:34
 * @Desc:
 */

type KVString struct {
    code    string
    message string
}

func NewKVString(code string, message string) KVString {
    return KVString{
        code:    code,
        message: message,
    }
}

func (kv KVString) Key() string {
    return kv.code
}

func (kv KVString) Value() string {
    return kv.message
}
