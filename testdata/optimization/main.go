package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/9/23 23:29
 * @Desc:
 */

func main() {
	// --- pprof 性能分析
	var isCpuPprof bool
	var isMemPprof bool
	flag.BoolVar(&isCpuPprof, "cpu", false, "是否开启cpu测试")
	flag.BoolVar(&isMemPprof, "mem", false, "是否开启内存测试")
	flag.Parse()
	
	if isCpuPprof {
		fmt.Println("开启cpu测试")
		file, err := os.Create("cpu.pprof")
		if err != nil {
			fmt.Println("create cpu.pprof file error:", err)
		}
		defer file.Close()
		pprof.StartCPUProfile(file) // 往文件中记录 CPU profile信息
		defer pprof.StopCPUProfile()
	} else {
		fmt.Println("不开启cpu测试")
	}
	
	fmt.Println("> 业务逻辑运行入口 开始运行")
	for i := 0; i < 2; i++ {
		fmt.Println("> 业务逻辑运行入口 ....", i)
		go PprofChannelReadPrint()
		go PprofChannelReadPrintV2()
		go PprofChannelReadPrintV3()
		go PprofChannelWriterReadPrint()
		go PprofChannelWriterReadPrintV2()
	}
	fmt.Println("> 业务逻辑运行入口 运行结束")
	time.Sleep(time.Second * 2)
	
	if isMemPprof {
		fmt.Println("开启内存测试")
		file, err := os.Create("mem.pprof")
		if err != nil {
			fmt.Println("create mem.pprof file error:", err)
		}
		defer file.Close()
		pprof.WriteHeapProfile(file) // 往文件中记录 内存 profile信息
		defer pprof.StopCPUProfile()
	} else {
		fmt.Println("不开启内存测试")
	}
	
	// --- pprof 性能分析
	
}

func PprofChannelReadPrint() {
	var streamInt chan int
	// streamInt = make(chan int, 10)
	for {
		select {
		case i, ok := <-streamInt:
			if !ok {
				fmt.Println("streamInt closed")
				return
			}
			fmt.Printf("streamInt: i:%d \n", i)
		default:
		}
	}
}

func PprofChannelReadPrintV2() {
	var streamInt chan int
	for {
		select {
		case i, ok := <-streamInt:
			if !ok {
				fmt.Println("streamInt closed")
				return
			}
			fmt.Printf("streamInt: i:%d \n", i)
		default:
			time.Sleep(time.Microsecond * 500) // 主动让出CPU，防止一直占用CPU
		}
	}
}

func PprofChannelReadPrintV3() {
	var streamInt chan int
	for {
		select {
		case i, ok := <-streamInt:
			if !ok {
				fmt.Println("streamInt closed")
				return
			}
			fmt.Printf("streamInt: i:%d \n", i)
		}
	}
}

func PprofChannelWriterReadPrint() {
	streamInt := make(chan int, 10)
	for i := 0; i < 20; i++ {
		go func(i int) {
			streamInt <- i
		}(i)
	}
	
	for {
		select {
		case i, ok := <-streamInt:
			if !ok {
				fmt.Println("streamInt closed")
				return
			}
			fmt.Printf("streamInt: i:%d \n", i)
		default:
		}
	}
}

func PprofChannelWriterReadPrintV2() {
	streamInt := make(chan int, 10)
	go func() {
		for {
			select {
			case i, ok := <-streamInt:
				if !ok {
					fmt.Println("streamInt closed")
					return
				}
				fmt.Printf("streamInt: i:%d \n", i)
			default:
			}
		}
	}()
	
	for i := 0; i < 20; i++ {
		streamInt <- i
	}
	close(streamInt)
}
