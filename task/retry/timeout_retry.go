package retry

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/10 20:48
 * @Desc:	用golang 的channel 实现对func taskFunc() error{} 方法的超时重传。
 *			这个重传的机制为：
 *				1. 如果正常收到响应则不需要重传。
 *				2. 如果长时间未收到响应则说明，taskFunc方法已经处理超时，需要重新执行taskFunc方法
 */

// 使用golang的select语句和time.After()函数来实现超时重传机制。具体来说，您可以创建一个channel，然后使用select语句监听该channel和time.After()函数返回的channel，如果在指定时间内没有收到响应，则会从time.After()函数返回的channel中读取数据，然后重新执行send方法。以下是一个示例代码：¹

func SendWithTimeout(timeout time.Duration, taskFunc func() error) error {
	ch := make(chan error)
	go func() {
		err := taskFunc()
		ch <- err
	}()
	select {
	case res := <-ch:
		return res
	case <-time.After(timeout):
		return SendWithTimeout(timeout, taskFunc)
	}
}
