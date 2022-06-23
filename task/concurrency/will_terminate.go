package concurrency

import (
    "context"
    "errors"
    "github.com/leeprince/goinfra/consts"
    "sync"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/12 下午11:05
 * @Desc:   并发执行
 */

type ConcurrencyTask struct {
    ctx          context.Context
    params       *params
    tasks        []taskFunc
    error        error
    errChan      chan struct{}
    taskDoneChan chan struct{}
}
type taskFunc func() error

type WillTerminateTask func()

func NewConcurrencyTask(ctx context.Context, opts ...WillTerminateOpt) *ConcurrencyTask {
    willTerminate := &ConcurrencyTask{
        ctx: ctx,
        params: &params{
            timeout:     defaultTimeout,
            stopOnError: false,
        },
        tasks:        nil,
        error:        nil,
        errChan:      make(chan struct{}, 1), // 有缓冲通信
        taskDoneChan: make(chan struct{}, 1), // 有缓冲通信
    }
    for _, opt := range opts {
        opt(willTerminate.params)
    }
    return willTerminate
}

func (t *ConcurrencyTask) AddTask(tasks ...taskFunc) {
    t.tasks = append(t.tasks, tasks...)
}

func (t *ConcurrencyTask) Start() error {
    if len(t.tasks) == 0 {
        return errors.New("err:len(t.tasks) == 0")
    }
    
    var doneTaskNum int
    doneTaskNumLock := sync.Mutex{}
    
    // wg := sync.WaitGroup{}
    for _, task := range t.tasks {
        go func(task taskFunc) {
            // defer wg.Done()
            taskErr := task()
            if taskErr != nil {
                // fmt.Println("......taskErr != nil:", taskErr)
                t.error = taskErr
                t.errChan <- struct{}{}
            }
            
            // 任务都执行完时，通过通道进行通信
            doneTaskNumLock.Lock()
            doneTaskNum++
            // fmt.Println(">>>>>>doneTaskNum:", doneTaskNum)
            doneTaskNumLock.Unlock()
            if doneTaskNum == len(t.tasks) {
                // fmt.Println(">>>>>>doneTaskNum done", doneTaskNum)
                t.taskDoneChan <- struct{}{}
            }
        }(task)
    }
    
    // 超时则返回超时错误
    tick := time.Tick(t.params.timeout)
    // fmt.Println("=========t.params.timeout > 0")
    for {
        select {
        case <-tick:
            return consts.TimeoutErr
        case <-t.taskDoneChan:
            // fmt.Println(">>>>>>case <-t.taskDoneChan")
            return t.error
        case <-t.errChan:
            if t.params.stopOnError {
                return t.error
            }
        }
    }
}
