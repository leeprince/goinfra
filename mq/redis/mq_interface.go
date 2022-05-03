package redis

import (
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/1 上午11:13
 * @Desc:   推送、拉取 redis 列表实现消息队列
 */

type MQ interface {
    Publish(key string, value interface{}) error
    Subscribe(key string, timetou time.Duration) (data []byte, err error)
}