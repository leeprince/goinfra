package kv

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/26 下午1:34
 * @Desc:
 */

type KStringVString struct {
    key     string
    message string
}

func NewKVString(code string, message string) KStringVString {
    return KStringVString{
        key:     code,
        message: message,
    }
}

func (kv KStringVString) Key() string {
    return kv.key
}

func (kv KStringVString) Value() string {
    return kv.message
}
