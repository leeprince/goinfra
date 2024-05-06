<!--

-->
# risk-common

`/risk-control/risk-common`，是风控系统各个微服务的公共包，包括`gostreaming`流式计算框架、`riskcontrol`业务公共变量等。

下图为风控系统的数据流处理，均依赖与`risk-common`包：
![风控系统](docs/images/风控系统.png)


## 介绍

### 项目结构

```bash
.
├── bin
├── cmd
│   ├── examples
│   │   └── word-count  # 流式计算框架示例程序 word count
├── docs
├── go.mod
├── go.sum
├── pkg
│   ├── dcachecli
│   ├── fileutil
│   ├── gostreaming     # 流式计算框架        
│   ├── grpchelper
│   ├── mysqlhelper
│   ├── redishelper
│   ├── riskcontrol     # 风控业务公共组件
│   ├── tars
│   ├── tarscli
│   └── transactions
└── README.md

```


## 流式计算框架 gostreaming 

### 介绍

gostreaming是一套无状态、轻量级的golang流式计算框架。之所以没有采用Flink/Spark，一方面是技术栈与当前后端组不太匹配，一方面是框架太重，部署麻烦，当前的业务体量还不需要，考虑到时间与学习成本故不采用。

三者区别：
- Flink: 有状态，真流式计算
- Spark：无状态，微批计算（伪流式）
- gostreaming：无状态，真流式计算

gostreaming与Flink类似，通过定义`DataSource`（数据源）与`DataStream`（数据流）来执行数据的转换、计算与聚合。无状态的优点在于方便水平伸缩，缺点在于中间计算状态依赖外部存储（Flink依赖与RocksDB与HDFS，gostreaming依赖于Redis）。

如果将流式计算视为一个数据操作的有向无环图，则`DataSource`与`DataStream`可以视为图中的节点，`DataSource`只生产数据/事件，`DataStream`接受数据/事件进行计算并传给下游的`DataStream`。

在gostreaming中，数据与事件是同一个概念，定义为：
```go
type Event struct {
	LogID string
	Data  interface{}
}
```
数据内容存放在`Event.Data`中。

### API

使用gostreaming编写一个实时 word count 的程序非常简单。


具体的完整代码可以查看`/risk-control/risk-common/cmd/examples/word-count`。


#### DataSource 数据源

##### 介绍

与Flink类似，gostreaming提供多种预先定义好的DataSource，比如`TCPDataSource`，`FileDataSource`。DataSource相当于数据的生产者，它的数据来源可以是上游消息队列，也可以是Redis Channel、MySQL数据库、文件内容，Socket等。

只需要实现了`DataSourceInterface`接口，或者匿名组合`*DataSource`结构体并覆盖`Start()`与`Stop()`方法，就可以作为gostreaming的数据生产器。

`DataSourceInterface`接口定义：
```go
type DataSourceInterface interface {
	Start()
	Stop()

	SetName(string)
	Print()
	Send(*Event)
	Delivery() <-chan *Event
}

```


##### 示例

`DataSource`结构体实现了`DataSourceInterface`接口的所有方法，但是`Start()`必须被覆盖。


```go
package main
// File: cmd/examples/word-count/main.go

import (
	"/risk-control/risk-common/cmd/examples/word-count/wordcount"
	"/risk-control/risk-common/pkg/gostreaming"
)

func main() {
	wordCountExample1()
}

func wordCountExample1() {
	// Engine 相当于一个脚手架，使用它来搭建流式计算有向无环图。
	engine := gostreaming.New(nil)

	// 使用gostreaming预定义的TCPDataSource，
	// 监听1379号端口，按行读取 socket 发送过来的文本内容，
	// 并以行为单位，构造Event发送给下游。
	dataSource := gostreaming.MustNewTCPDataSource(1379)
	dataSource.SetName("word-count-example-1") // optional，设置节点名称。
	dataSource.Print()                         // optional，打印 Event.Data

	// 设置 DataSource
	engine.SetDataSource(dataSource)

	defer engine.Stop()
	engine.Run()
}
```

通过`go run main.go`启动服务后，在Linux下可以使用`nc`命令行程序发送文本行内容：
```bash
nc localhost 1379
```


然后输入两行内容：
```txt
hello world
hello golden-cloud
```

这时，word-count程序输出：
```txt
[gostreaming] [INFO] [DataSource:word-count-example-1]: event.Data: hello world
[gostreaming] [INFO] [DataSource:word-count-example-1]: event.Data: hello golden-cloud
```

可以看见，`TCPDataSource`产生了两个事件，数据内容分别为字符串`hello world`与`hello golden-cloud`。


#### DataStream 数据流

##### 介绍

和`DataSource`与`DataSourceInterface`的关系类似，`DataStream`结构体实现了`DataStreamInterface`接口的所有方法：

```go
type DataStreamInterface interface {
	Process(StatusStorage, <-chan *Event)

	SetName(string)
	Name() string

	SetParallel(int)
	Parallel() int

	Send(*Event)
	Print()
	Delivery() <-chan *Event
}

```

自定义的`DataStream`只需要匿名组合`*DataStream`并覆盖`Process(StatusStorage, <-chan *Event)`方法即可。

##### 示例

继续上面的 word count 例子，定义一个`WordSpliter`进行分词：
```go
package wordcount
// File: cmd/examples/word-count/wordcount/word_spliter.go

import (
	"strings"
	"/risk-control/risk-common/pkg/gostreaming"
)

var _ gostreaming.DataStreamInterface = (*WordSpliter)(nil)

type WordSpliter struct {
	*gostreaming.DataStream
}

func NewWordSpliter() gostreaming.DataStreamInterface {
	return &WordSpliter{
		DataStream: gostreaming.NewDataStream(),
	}
}

func (w *WordSpliter) Process(_ gostreaming.StatusStorage, ch <-chan *gostreaming.Event) {
	for {
		select {
		case event := <-ch:
			line := event.Data.(string)
			words := strings.Split(line, " ")

			for _, word := range words {
				word = strings.TrimSpace(word)
				if word == "" {
					continue
				}
				w.Send(&gostreaming.Event{
					Data: word,
				})
			}
		}
	}
}

```


然后修改我们的`main.go`：

```go
package main
// File: cmd/examples/word-count/main.go


import (
	"/risk-control/risk-common/cmd/examples/word-count/wordcount"
	"/risk-control/risk-common/pkg/gostreaming"
)

func main() {
	// wordCountExample1()
	wordCountExample2()
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
```


重新编译启动程序，然后同样继续使用`nc`发送同样的内容：

```bash
nc localhost 1379
hello world
hello golden-cloud
```

这时程序输出为：
```
[gostreaming] [INFO] [DataSource:word-count-example-1]: event.Data: hello world
[gostreaming] [INFO] [DataStream:word-spliter]: event.Data: hello
[gostreaming] [INFO] [DataStream:word-spliter]: event.Data: world
[gostreaming] [INFO] [DataSource:word-count-example-1]: event.Data: hello golden-cloud
[gostreaming] [INFO] [DataStream:word-spliter]: event.Data: hello
[gostreaming] [INFO] [DataStream:word-spliter]: event.Data: golden-cloud
```

可以看见，`word-spliter`这个`DataStream`将文本行分词后，以一个词构造一个`Event`，共发送4个`Event`给下游。



#### StatusStorage 状态存储器

##### 介绍


在gostreaming中，中间计算状态可以存放在程序的全局变量中（比如`sync.Map`中），也可以存放在外部存储（例如Redis、MySQL、LevelDB等）。gostreaming将状态存储抽象为一套接口：`StatusStorage`，通过实现了`BatchInterface`接口的`Batch`来实现批量读/写：

```go
package gostreaming
// File: pkg/gostreaming/status_sotrage.go

type StatusStorage interface {
	ExecBatch(BatchInterface) ([]interface{}, []error, error)
}

```

gostreaming预先封装好了两种状态存储器：`MemoryStatusStorage`与`RedisStatusStorage`，分别使用内存和Redis作为具体存储实现。

`Batch`提供的状态读写接口与Redis提供的功能类似：
```go
package gostreaming
// File: pkg/gostreaming/batch.go

func NewBatch() *Batch 
func (b *Batch) Size() int 
func (b *Batch) GetBatchCommands() []*BatchCommand 
func (b *Batch) String() string 

func (b *Batch) Get(primaryKeys []string, targetName string, descriptions []string) 
func (b *Batch) SCard(primaryKeys []string, targetName string, descriptions []string) 
func (b *Batch) SMembers(primaryKeys []string, targetName string, descriptions []string) 
func (b *Batch) HLen(primaryKeys []string, targetName string, descriptions []string) 
func (b *Batch) Incr(primaryKeys []string, targetName string, descriptions []string) 
func (b *Batch) Set(primaryKeys []string, targetName string, descriptions []string, value interface{}) 
func (b *Batch) SAdd(primaryKeys []string, targetName string, descriptions []string, value interface{}) 
func (b *Batch) HIncrBy(primaryKeys []string, targetName string, descriptions []string, field string, incrBy int) 
func (b *Batch) HSet(primaryKeys []string, targetName string, descriptions []string, field string, value interface{}) 
func (b *Batch) HGet(primaryKeys []string, targetName string, descriptions []string, field string) 
func (b *Batch) ZAdd(primaryKeys []string, targetName string, descriptions []string, score float64, member interface{}) 
func (b *Batch) ZRangeByScoreWithScores(primaryKeys []string, targetName string, descriptions []string, min string, max string, offset int64, count int64) 
func (b *Batch) ZRevRangeByScoreWithScores(primaryKeys []string, targetName string, descriptions []string, max string, min string, offset int64, count int64) 
func (b *Batch) ZCount(primaryKeys []string, targetName string, descriptions []string, min string, max string) 


```

##### 示例

完成分词的实现后，接下来实现单词计数的`DataStream`：

```go
package wordcount
// File: cmd/examples/word-count/wordcount/word_counter.go

import (
	"/risk-control/risk-common/pkg/gostreaming"
)

var _ gostreaming.DataStreamInterface = (*WordCounter)(nil)

type WordCounter struct {
	*gostreaming.DataStream
}

func NewWordCounter() gostreaming.DataStreamInterface {
	return &WordCounter{
		DataStream: gostreaming.NewDataStream(),
	}
}

func (w *WordCounter) Process(statusStorage gostreaming.StatusStorage, ch <-chan *gostreaming.Event) {
	for {
		select {
		// 监听事件
		case event := <-ch:
			word := event.Data.(string)
			batch := gostreaming.NewBatch()

			batch.Incr([]string{word}, "word-count", nil)
			statusStorage.ExecBatch(batch)

			w.Send(event)
		}
	}
}

```

然后再实现一个每隔1秒打印所有单词计数的`WordCountPrinter`：


```go
package wordcount
// File: cmd/examples/word-count/wordcount/word_count_printer.go

import (
	"fmt"
	"time"

	"github.com/emirpasic/gods/sets/treeset"
	"/risk-control/risk-common/pkg/gostreaming"
)

var _ gostreaming.DataStreamInterface = (*WordCountPrinter)(nil)

type WordCountPrinter struct {
	*gostreaming.DataStream
}

func NewWordCountPrinter() gostreaming.DataStreamInterface {
	return &WordCountPrinter{
		DataStream: gostreaming.NewDataStream(),
	}
}

func (w *WordCountPrinter) Process(statusStorage gostreaming.StatusStorage, ch <-chan *gostreaming.Event) {
	ticker := time.NewTicker(1 * time.Second)
	wordSet := treeset.NewWithStringComparator()
	iterTime := 0

	for {
		select {
		case event := <-ch:
			word := event.Data.(string)
			wordSet.Add(word)
			w.Send(event)
		case <-ticker.C:
			words, counts := w.getWordCounts(statusStorage, wordSet)
			if len(words) > 0 {
				fmt.Println("")
			}
			for i := range words {
				fmt.Printf("[INFO] [word-count-printer] [%d] %s: %d \n", iterTime, words[i], counts[i])
			}

			iterTime++

		}
	}
}

func (w *WordCountPrinter) getWordCounts(statusStorage gostreaming.StatusStorage, wordSet *treeset.Set) ([]string, []int) {
	batch := gostreaming.NewBatch()

	words := make([]string, wordSet.Size())
	counts := make([]int, wordSet.Size())

	for i, wordInterface := range wordSet.Values() {
		word := wordInterface.(string)
		words[i] = word
		batch.Get([]string{word}, "word-count", nil)
	}

	results, errs, err := statusStorage.ExecBatch(batch)
	if err != nil {
		fmt.Printf("[ERROR] err: %+v \n", err)
		return nil, nil
	}

	for _, err := range errs {
		if err != nil {
			fmt.Printf("[WARN] err: %+v \n", err)
		}
	}

	for i, result := range results {
		counts[i] = result.(int)
	}
	return words, counts
}


```


最后，将`WordCount`与`WordCountPrinter`两个`DataStream`加入`engine`中：


```go
package main
// File: cmd/examples/word-count/main.go

import (
	"/risk-control/risk-common/cmd/examples/word-count/wordcount"
	"/risk-control/risk-common/pkg/gostreaming"
)

func main() {
	wordCountExample3()
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

```


继续使用`nc`来测试：

```bash
nc localhost 1379
hello world
hello golden-cloud
```


可以看见程序每秒打印所有单词的计数输出为：
```txt
[INFO] [word-count-printer] [7] hello: 1 
[INFO] [word-count-printer] [7] world: 1 

[INFO] [word-count-printer] [8] golden-cloud: 1 
[INFO] [word-count-printer] [8] hello: 2 
[INFO] [word-count-printer] [8] world: 1 

[INFO] [word-count-printer] [9] golden-cloud: 1 
[INFO] [word-count-printer] [9] hello: 2 
[INFO] [word-count-printer] [9] world: 1 

[INFO] [word-count-printer] [10] golden-cloud: 1 
[INFO] [word-count-printer] [10] hello: 2 
[INFO] [word-count-printer] [10] world: 1 
```