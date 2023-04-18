package main

import (
	"fmt"
	"log"
	"runtime"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/4 10:36
 * @Desc:
 */
func main() {
	test1()
}

func test1() {
	test2()
}

func test2() {
	test3()
}

func test3() {
	//fmt.Println("==============================CallerFunc()==============================")
	//CallerFunc()
	//fmt.Println("==============================CallerFuncAll()==============================")
	//CallerFuncAll()

	fmt.Println("==============================CallersFunc()==============================")
	CallersFunc()
	fmt.Println("==============================CallersFuncAll()==============================")
	CallersFuncAll()
}

func CallerFunc() {
	//CallerFunc0()
	//CallerFunc1()
	//CallerFunc2()
	CallerFunc3()
}

func CallerFuncAll() {
	CallerFunc0()
	CallerFunc1()
	CallerFunc2()
	CallerFunc3()
}

func CallerFunc0() {
	fmt.Println("--------------------------CallerFunc 0")
	pc, file, line, ok := runtime.Caller(0)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f := runtime.FuncForPC(pc)
	log.Println(f.Name())
}
func CallerFunc1() {
	fmt.Println("--------------------------CallerFunc 1")
	pc, file, line, ok := runtime.Caller(1)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f := runtime.FuncForPC(pc)
	log.Println(f.Name())
}
func CallerFunc2() {
	fmt.Println("--------------------------CallerFunc 2")
	pc, file, line, ok := runtime.Caller(2)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f := runtime.FuncForPC(pc)
	log.Println(f.Name())
}
func CallerFunc3() {
	fmt.Println("--------------------------CallerFunc 3")
	pc, file, line, ok := runtime.Caller(3)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f := runtime.FuncForPC(pc)
	log.Println(f.Name())
}

func CallersFunc() {
	// 以下需要逐个解开注释，仅允许单个运行测试才准确

	//CallersFunc0()
	//CallersFunc1()
	//CallersFunc2()
	CallersFunc3()
}

func CallersFuncAll() {
	// 解开所有注释

	CallersFunc0()
	CallersFunc1()
	CallersFunc2()
	CallersFunc3()
}

func CallersFunc0() {
	pc := make([]uintptr, 1)

	fmt.Println("--------------------------CallersFunc 0")
	runtime.Callers(0, pc)
	f := runtime.FuncForPC(pc[0])
	log.Println(f)
	log.Println(f.Name())
}

func CallersFunc1() {
	pc := make([]uintptr, 1)

	fmt.Println("--------------------------CallersFunc 1")
	runtime.Callers(1, pc)
	f := runtime.FuncForPC(pc[0])
	log.Println(f)
	log.Println(f.Name())
}

func CallersFunc2() {
	pc := make([]uintptr, 1)

	fmt.Println("--------------------------CallersFunc 2")
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	log.Println(f)
	log.Println(f.Name())
}

func CallersFunc3() {
	pc := make([]uintptr, 1)

	fmt.Println("--------------------------CallersFunc 3")
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	log.Println(f)
	log.Println(f.Name())
}
