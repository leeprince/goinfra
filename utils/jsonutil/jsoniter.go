package jsonutil

import jsoniter "github.com/json-iterator/go"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/28 18:15
 * @Desc:	jsoniter 与标准库中的json 类似，但是它的性能更好。
 */

var JsoniterCompatible = jsoniter.ConfigCompatibleWithStandardLibrary
