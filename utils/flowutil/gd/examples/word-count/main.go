package main

/*
 * @Date: 2020-09-03 16:50:34
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-09-04 16:55:38
 */

import (
	"gitlab.yewifi.com/risk-control/risk-common/cmd/examples/word-count/wordcount"
	"gitlab.yewifi.com/risk-control/risk-common/pkg/gostreaming"
)

func main() {
	// wordCountExample1()
	// wordCountExample2()
	wordCountExample3()
}

func wordCountExample1() {
	// Engine 相当于一个脚手架，使用它来搭建流式计算有向无环图。
	engine := gostreaming.New(nil)

	// 使用gostreaming预定义的TCPDataSource，
	// 监听1370号端口，按行读取 socket 发送过来的文本内容，
	// 并以行为单位，构造Event发送给下游。
	dataSource := gostreaming.MustNewTCPDataSource(1379)
	dataSource.SetName("word-count-example-1") // optional，设置节点名称。
	dataSource.Print()                         // optional，打印 Event.Data

	// 设置 DataSource
	engine.SetDataSource(dataSource)

	defer engine.Stop()
	engine.Run()
}

func wordCountExample2() {
	engine := gostreaming.New(nil)

	dataSource := gostreaming.MustNewTCPDataSource(1379)
	dataSource.SetName("word-count-example-1")
	dataSource.Print()

	// 创建一个DatStream，负责将文本行按空格进行分词。
	wordSpliter := wordcount.NewWordSpliter()
	wordSpliter.SetName("word-spliter")
	wordSpliter.Print()
	wordSpliter.SetParallel(10) // 设置并发协程数，默认为1

	engine.
		SetDataSource(dataSource).
		AddDataStream(wordSpliter)

	defer engine.Stop()
	engine.Run()
}

func wordCountExample3() {
	statusStorage := gostreaming.NewMemoryStatusStorage()
	engine := gostreaming.New(statusStorage)

	engine.
		SetDataSource(gostreaming.MustNewTCPDataSource(1379)).
		AddDataStream(wordcount.NewWordSpliter()).
		AddDataStream(wordcount.NewWordCounter()).
		AddDataStream(wordcount.NewWordCountPrinter())

	defer engine.Stop()
	engine.Run()
}
