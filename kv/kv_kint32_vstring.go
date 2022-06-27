package kv

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/26 下午1:34
 * @Desc:   键值对管理
 */

type KInt32VString struct {
    key     int32
    message string
}

func NewKInt32VString(code int32, message string) KInt32VString {
    return KInt32VString{
        key:     code,
        message: message,
    }
}

func (kv KInt32VString) Key() int32 {
    return kv.key
}

func (kv KInt32VString) Value() string {
    return kv.message
}
