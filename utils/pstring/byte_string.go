package pstring

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/13 18:14
 * @Desc:	字符串与字节互转
 */

import "unsafe"

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
