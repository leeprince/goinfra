package concurrency

import (
    "context"
    "errors"
    "fmt"
    "github.com/leeprince/goinfra/utils/ptime"
    "golang.org/x/sync/errgroup"
    "sync"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/14 下午10:51
 * @Desc:
 */

// 测试出现竞态资源问题
//  竞态资源问题：本质是并发访问或更新同一个变量
//      1.并发更新同一个变量：
//          1)append：在协程环境下可能出现错误,当A和B两个协程运行append的时候同时发现s[1]这个位置是空的，他们就都会把自己的值放在这个位置，这样他们两个的值就会覆盖，造成数据丢失，导致竞态资源问题。
//              - 解决：
//                  a.加互斥锁更新【推荐：简单、稳定、高效率；】
//                  b.使用通道更新【优点：灵活】：注意一个问题：读通道之后，可能会出现立即结束所有循环计算的协程的情况，那么会导致 taskNum 还没计算完最后一个数字就响应了，导致计算的总数少。解决：所有计算的协程结束后添加个很小的等待时间 `time.Sleep(time.Nanosecond)`
//          2)循环（循环条件与循环代码的变量是不共享的）的循环代码中存在多个协程并发更新同一个变量
//              - 解决：
//                  a.加互斥锁更新【推荐】
//                  b.使用通道更新：注意一个问题：读通道之后，可能会出现立即结束所有循环计算的协程的情况，那么会导致 taskNum 还没计算完最后一个数字就响应了，导致计算的总数少。解决：所有计算的协程结束后添加个很小的等待时间 `time.Sleep(time.Nanosecond)`
//      2.并发访问同一个变量
//          1）循环（循环条件与循环代码的变量是不共享的）的循环代码中存在多个协程并发访问循环条件的同一个变量
//              a.循环切片的循环代码中存在多个协程并发访问循环条件的同一个变量，由于是同一个变量，导致协程访问到的变量可能是经过并发多次循环后的当时变量导致竞态资源问题。
//                  - 解决：具体说明：https://golang.org/doc/faq#closures_and_goroutines（需翻墙）。
//                      a)使用临时变量：利用 `循环条件与循环代码的变量是不共享的` 的特性，在循环代码中重新将循环条件的变量赋值到新的变量作为临时变量，并临时变量代替循环条件的变量。
//                      b)绑定参数：直接将循环条件的变量绑定到使用协程的闭包函数的参数中进行访问
func TestConcurrency1(t *testing.T) {
    // --- 出现竞态资源问题:1.并发更新同一个变量：1)append的情况 ----------
    mtime := ptime.NewTimeCost("TestConcurrency1")
    iMax := 100000
    
    wg := sync.WaitGroup{}
    s := make([]int, 0)
    for i := 0; i < iMax; i++ {
        wg.Add(1)
        go func() {
            s = append(s, i)
            wg.Done()
        }()
    }
    wg.Wait()
    mtime.Duration("TestConcurrency1 s")
    fmt.Printf("》》》s 出现竞态资源问题:1.并发更新同一个变量：1)append的情况： %v\n", len(s))
    
    // 出现竞态资源问题:1.并发更新同一个变量：1)append的情况
    //      解决：a.加互斥锁更新
    mtime.Duration("TestConcurrency1 减去之前打印时间")
    wg = sync.WaitGroup{}
    mutex := sync.Mutex{}
    s1 := make([]int, 0)
    for i := 0; i < iMax; i++ {
        wg.Add(1)
        go func() {
            mutex.Lock()
            s1 = append(s1, i)
            mutex.Unlock()
            wg.Done()
        }()
    }
    wg.Wait()
    mtime.Duration("TestConcurrency1 s1")
    fmt.Printf("》》》s1 s 出现竞态资源问题:1.并发更新同一个变量：1)append的情况 》解决：a.加互斥锁更新： %v\n", len(s1))
    
    // 出现竞态资源问题:1.并发更新同一个变量：1)append的情况
    //      解决：b.使用通道更新
    mtime.Duration("TestConcurrency1 减去之前打印时间")
    wg = sync.WaitGroup{}
    ch := make(chan int)
    s2 := make([]int, 0)
    for i := 0; i < iMax; i++ {
        wg.Add(1)
        // 写入通道
        go func(i int) {
            ch <- i
            wg.Done()
        }(i)
        
    }
    // 读出通道方式一【推荐】
    go func() {
        for {
            // i, ok := <-ch
            i := <-ch
            s2 = append(s2, i)
        }
    }()
    // 读出通道方式二 // 此处只有一个通道等待，推荐使用方式一
    // 消费
    // go func() {
    //     for {
    //         select {
    //         case i := <-ch:
    //             s2 = append(s2, i)
    //         }
    //     }
    // }()
    wg.Wait()
    mtime.Duration("TestConcurrency1 s2")
    fmt.Printf("》》》s2 s 出现竞态资源问题:1.并发更新同一个变量：1)append的情况 》解决：b.使用通道更新： %v\n", len(s2))
    // --- 出现竞态资源问题:1.并发更新同一个变量：1)append的情况 -end ----------
}

func TestConcurrency2(t *testing.T) {
    // --- 出现竞态资源问题:1.并发更新同一个变量：1)循环的情况 ----------
    iMax := 10000
    wg := sync.WaitGroup{}
    
    mtime := ptime.NewTimeCost("TestConcurrency1")
    var taskNum int
    task1 := func() {
        for i := iMax; i > 0; i-- {
            taskNum++
        }
    }
    task2 := func() {
        for i := iMax; i > 0; i-- {
            taskNum++
        }
    }
    
    wg.Add(1)
    go func() {
        task1()
        wg.Done()
    }()
    
    wg.Add(1)
    go func() {
        task2()
        wg.Done()
    }()
    wg.Wait()
    mtime.Duration("-")
    fmt.Println(">>> 出现竞态资源问题:1.并发更新同一个变量：1)循环的情况: taskNum", taskNum)
    
    // 出现竞态资源问题:1.并发更新同一个变量：1)循环的情况
    //    解决：a.加互斥锁更新
    mtime.Duration("TestConcurrency2 减去之前打印时间")
    wg = sync.WaitGroup{}
    lk := sync.Mutex{}
    taskNum = 0
    task1 = func() {
        for i := iMax; i > 0; i-- {
            lk.Lock()
            taskNum++
            lk.Unlock()
        }
    }
    task2 = func() {
        for i := iMax; i > 0; i-- {
            lk.Lock()
            taskNum++
            lk.Unlock()
        }
    }
    
    wg.Add(1)
    go func() {
        task1()
        wg.Done()
    }()
    
    wg.Add(1)
    go func() {
        task2()
        wg.Done()
    }()
    wg.Wait()
    mtime.Duration("-")
    fmt.Println(">>> 出现竞态资源问题:1.并发更新同一个变量：1)循环的情况 》 解决：a.加互斥锁更新 taskNum", taskNum)
    
    // 出现竞态资源问题:1.并发更新同一个变量：1)循环的情况
    //    解决：b.使用通道更新
    mtime.Duration("TestConcurrency2 减去之前打印时间")
    wg = sync.WaitGroup{}
    numCh := make(chan int)
    taskNum = 0
    task1 = func() {
        for i := iMax; i > 0; i-- {
            numCh <- 1
        }
    }
    task2 = func() {
        for i := iMax; i > 0; i-- {
            numCh <- 1
        }
    }
    
    wg.Add(1)
    go func() {
        task1()
        wg.Done()
    }()
    
    wg.Add(1)
    go func() {
        task2()
        wg.Done()
    }()
    
    // 读通道
    //  注意一个问题：读通道之后，可能会出现立即结束所有循环计算的协程的情况，那么会导致 taskNum 还没计算完最后一个数字就响应了，导致计算的总数少。
    //      解决：
    //          - 如果等待协程结束后，需要立即读出该竞态资源变量，则需要在等待协程结束后添加个很小很小的等待时间 `time.Sleep(time.Nanosecond)`
    //          - 如果等待协程结束后，不需要立即读出该竞态资源变量，而是经过其他程序计算再读取，则完全可以忽略这个问题，即不需要添加这个很小很小的等待时间
    // 消费
    go func() {
        for {
            // i, ok := <-numCh
            <-numCh
            taskNum += 1
        }
    }()
    wg.Wait()
    fmt.Println("taskNum：", taskNum)
    time.Sleep(time.Nanosecond) // 很小很小的等待时间
    fmt.Println("taskNum：", taskNum)
    
    // 用于统计时间及
    mtime.Duration("-")
    fmt.Println(">>> 出现竞态资源问题:1.并发更新同一个变量：1)循环的情况 》 解决：b.使用通道更新 taskNum", taskNum)
    
    // --- 出现竞态资源问题:1.并发更新同一个变量：1)循环的情况 -end ----------
}
func TestConcurrency3(t *testing.T) {
    // --- 出现竞态资源问题:2.并发访问同一个变量 1）循环（循环条件与循环代码的变量是不共享的）的循环代码中存在多个协程并发访问循环条件的同一个变量 -----
    numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    wg := sync.WaitGroup{}
    
    wg.Add(len(numSlice))
    for _, i2 := range numSlice {
        go func() {
            fmt.Println("出现竞态资源问题:", i2)
            wg.Done()
        }()
    }
    wg.Wait()
    
    // 解决：使用临时变量
    wg.Add(len(numSlice))
    for _, i2 := range numSlice {
        i2 := i2
        go func() {
            fmt.Println("出现竞态资源问题 》解决：使用临时变量:", i2)
            wg.Done()
        }()
    }
    wg.Wait()
    
    // 解决：绑定参数
    wg.Add(len(numSlice))
    for _, i2 := range numSlice {
        go func(i2 int) {
            fmt.Println("出现竞态资源问题 》解决：绑定参数:", i2)
            wg.Done()
        }(i2)
    }
    wg.Wait()
    
    // --- 出现竞态资源问题:2.并发访问同一个变量 1）循环（循环条件与循环代码的变量是不共享的）的循环代码中存在多个协程并发访问循环条件的同一个变量 -end -----
}

// 测试竞态资源问题
func TestWillTerminate(t *testing.T) {
    var task1Num int
    task1 := func() {
        for i := 10000; i > 0; i-- {
            task1Num++
        }
    }
    var task2Num int
    task2 := func() {
        for i := 20000; i > 0; i-- {
            task2Num++
        }
    }
    
    task1()
    task2()
    fmt.Println("--- task1Num task2Num:")
    fmt.Println("task1Num:", task1Num)
    fmt.Println("task2Num:", task2Num)
    
    var task3Num int
    task3 := func() {
        for i := 100000; i > 0; i-- {
            task3Num++
        }
    }
    var task4Num int
    task4 := func() {
        for i := 200000; i > 0; i-- {
            task4Num++
        }
    }
    
    go task3()
    go task4()
    time.Sleep(time.Second * 1)
    fmt.Println("--- task3Num task4Num:")
    fmt.Println("task3Num:", task3Num)
    fmt.Println("task4Num:", task4Num)
    
    // 出现竞态资源问题
    var task5Num int
    task5_1 := func() {
        for i := 100000; i > 0; i-- {
            task5Num++
        }
    }
    task5_2 := func() {
        for i := 200000; i > 0; i-- {
            task5Num++
        }
    }
    go task5_1()
    go task5_2()
    time.Sleep(time.Second * 1)
    fmt.Println("--- task5Num:")
    fmt.Println("task5Num:", task5Num)
    
    // 出现竞态资源问题：加互斥锁解决
    var task6Num int
    l := sync.Mutex{}
    task6_1 := func() {
        for i := 100000; i > 0; i-- {
            l.Lock()
            task6Num++
            l.Unlock()
        }
    }
    task6_2 := func() {
        for i := 200000; i > 0; i-- {
            l.Lock()
            task6Num++
            l.Unlock()
        }
    }
    go task6_1()
    go task6_2()
    time.Sleep(time.Second * 1)
    fmt.Println(">>> task6Num:")
    fmt.Println("task6Num:", task6Num)
    
    // 出现竞态资源问题：加互斥锁解决
    var task7Num int
    lk := sync.Mutex{}
    task7_1 := func() {
        for i := 100000; i > 0; i-- {
            lk.Lock()
            task7Num++
            lk.Unlock()
        }
    }
    task7_2 := func() {
        for i := 200000; i > 0; i-- {
            lk.Lock()
            task7Num++
            lk.Unlock()
        }
    }
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() {
        defer wg.Done()
        task7_1()
    }()
    go func() {
        defer wg.Done()
        task7_2()
    }()
    wg.Wait()
    // time.Sleep(time.Second * 1)
    fmt.Println(">>> task7Num:")
    fmt.Println("task7Num:", task7Num)
    
    // 出现竞态资源问题：加互斥锁解决
    var task8Num int
    lk = sync.Mutex{}
    task8_1 := func() {
        for i := 100000; i > 0; i-- {
            lk.Lock()
            task8Num++
            lk.Unlock()
        }
    }
    task8_2 := func() error {
        for i := 200000; i > 0; i-- {
            lk.Lock()
            task8Num++
            lk.Unlock()
        }
        
        // 测试另外一个协程发生错误时，是否仍需要等待当前协程执行完毕? 答案：是需要的！
        // time.Sleep(time.Second * 2)
        // fmt.Println("task8_2 测试另外一个协程发生错误时，是否仍需要等待当前协程执行完毕?")
        return nil
        // return errors.New(">>>>>task8_2 test error")
    }
    errgp := errgroup.Group{}
    errgp.Go(func() error {
        task8_1()
        
        // 测试另外一个协程发生错误时，是否仍需要等待当前协程执行完毕? 答案：是需要的！
        // time.Sleep(time.Second * 2)
        // fmt.Println("task8_1 测试另外一个协程发生错误时，是否仍需要等待当前协程执行完毕?")
        return nil
        // return errors.New(">>>>>task8_1 test error")
    })
    errgp.Go(task8_2)
    err := errgp.Wait()
    if err != nil {
        fmt.Println("errgp.Wait() err:", err)
    }
    // time.Sleep(time.Second * 1)
    fmt.Println(">>> task8Num:")
    fmt.Println("task8Num:", task8Num)
}

func TestNewWillTerminate(t *testing.T) {
    type args struct {
        ctx  context.Context
        opts []WillTerminateOpt
    }
    tests := []struct {
        name string
        args args
        want *ConcurrencyTask
    }{
        {
            args: args{
                ctx: context.Background(),
                opts: []WillTerminateOpt{
                    WithTimeout(time.Second * 1),
                    // WithStopOnError(false),
                    WithStopOnError(true),
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            task := NewConcurrencyTask(tt.args.ctx, tt.args.opts...)
            
            var taskNum int
            lk := sync.Mutex{}
            task1 := func() error {
                for i := 100000; i > 0; i-- {
                    lk.Lock()
                    taskNum++
                    lk.Unlock()
                }
                
                // 测试另外一个协程发生错误时，是否仍需要等待当前协程执行完毕?
                time.Sleep(time.Second * 2)
                // fmt.Println(">>> task1")
                // return nil
                return errors.New(">>> task1 测试另外一个协程发生错误时，是否仍需要等待当前协程执行完毕?")
            }
            task2 := func() error {
                for i := 200000; i > 0; i-- {
                    lk.Lock()
                    taskNum++
                    // time.Sleep(time.Second * 2) // 测试超时，并且发生错误时任务终止的情况是否影响结果
                    lk.Unlock()
                }
                
                // 测试另外一个协程发生错误时，是否仍需要等待当前协程执行完毕?
                time.Sleep(time.Second * 2)
                // fmt.Println(">>> task2")
                return nil
                // return errors.New(">>> task2 测试另外一个协程发生错误时，是否仍需要等待当前协程执行完毕?")
            }
            
            task.AddTask(task1, task2)
            
            err := task.Start()
            fmt.Println()
            fmt.Println("err:", err)
            fmt.Println("taskNum:", taskNum)
        })
    }
}

