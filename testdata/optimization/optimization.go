package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/9/23 23:38
 * @Desc:
 */

func SliceStr(str, seq string) []string {
	if len(str) <= 0 {
		return []string{}
	}
	index := strings.Index(str, seq)
	
	/*
		定义 var ret []string 时基准测试为结果为如下
		➜  testdata git:(master) ✗ go test -v  -bench=BenchmarkSliceStr -run=none -benchmem
		goos: darwin
		goarch: amd64
		pkg: github.com/leeprince/goinfra/testdata
		cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
		BenchmarkSliceStr
		BenchmarkSliceStr-12             5189738               225.1 ns/op           112 B/op          3 allocs/op
		---
		基准测试了 5189738 次，每次 耗时 225.1 ns，每次分配 3 次内存，每次分配 112 B。
	
		针对每次分配 3 次内存可以进行优化,因为通过代码我们可以知道 ret = append(ret, str[:index])在 对切片扩容的情况下会进行一次底层数组的修改，并且会发生内存拷贝，所以我们可以使用 make() 函数来创建一个指定长度和容量的切片，避免底层数组拷贝。
	
		为了减少内存分配次数，定义ret修改为 ret :=make([]string, 0, strings.Count(str, seq)+1)
		继续进行基准测试发现，内存分配次数降下来了，并且可以跑的基准测试次数变多了，每次耗时更好，分配内存更少。
		➜  testdata git:(master) ✗ go test -v  -bench=BenchmarkSliceStr -run=none -benchmem
		goos: darwin
		goarch: amd64
		pkg: github.com/leeprince/goinfra/testdata
		cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
		BenchmarkSliceStr
		BenchmarkSliceStr-12            10211023               115.6 ns/op            48 B/op          1 allocs/op
		PASS
		ok      github.com/leeprince/goinfra/testdata   1.684s
	
	*/
	ret := make([]string, 0, strings.Count(str, seq)+1)
	
	for index > 0 {
		ret = append(ret, str[:index])
		str = str[index+len(seq):]
		index = strings.Index(str, seq)
	}
	if strings.Trim(str, " ") != "" {
		ret = append(ret, str)
	}
	if index > 0 {
		fmt.Println("index > 0，无效代码，仅为测试代码覆盖率")
	}
	return ret
}
