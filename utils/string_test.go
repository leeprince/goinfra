package utils

import (
    "fmt"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午10:10
 * @Desc:
 */

func TestRune(t *testing.T) {
    // msg := "leeprince"
    msg := "我爱您，中国"
    messageRunes := []rune(msg)
    fmt.Println(messageRunes)
    fmt.Println(string(messageRunes))
    
    messagebytes := []byte(msg)
    fmt.Println(messagebytes)
    fmt.Println(string(messagebytes))
}