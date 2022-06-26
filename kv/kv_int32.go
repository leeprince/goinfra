package kv

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/26 下午1:34
 * @Desc:
 */

type KVInt32 struct {
    code    int32
    message string
}

func NewKeyValue(code int32, message string) KVInt32 {
    return KVInt32{
        code:    code,
        message: message,
    }
}

func (kv KVInt32) Key() int32 {
    return kv.code
}

func (kv KVInt32) Value() string {
    return kv.message
}
