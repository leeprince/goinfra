package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/3 14:28
 * @Desc:
 */

import (
	"fmt"
	"github.com/speps/go-hashids"
)

func main() {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	hd.Alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	h, _ := hashids.NewWithData(hd)
	
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
	
	fmt.Println("---")
	
	e, _ = h.Encode([]int{45})
	fmt.Println(e)
	d, _ = h.DecodeWithError(e)
	fmt.Println(d)
	fmt.Println(">>> ----")
	
	e, _ = h.EncodeInt64([]int64{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ = h.DecodeWithError(e)
	fmt.Println(d)
	
	e, _ = h.EncodeInt64([]int64{45})
	fmt.Println(e)
	d, _ = h.DecodeWithError(e)
	fmt.Println(d)
}
