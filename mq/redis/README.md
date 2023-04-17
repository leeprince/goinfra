# redis 实现消息队列

> 文件结构

```
├── delay_mq_interface.go // 延迟队列接口
├── delay_mq_zadd_zrangebyscore_zrem_imp.go // zadd + zrangebyscore&zrem 实现延迟队列接口
├── delay_mq_zadd_zrangebyscore_zrem_imp_test.go // 延迟队列接口实现测试
├── mq_interface.go // 普通消息队列接口
├── mq_push_pop_imp.go // lpush + rpop (默认，左近右出)|| rpush + lpop 实现普通消息队列接口 
├── mq_push_pop_imp_test.go // 普通消息队列接口实现测试
├── pubsub_mq_interface.go  // 发布订阅者模型接口
└── pubsub_mq_publish_subscribe_imp.go // `publish`、`subscribe` 发布订阅命令实现 PubSubMQ 接口
└── pubsub_mq_publish_subscribe_imp_test.go // 发布订阅者模型接口实现测试
```

## 普通消息队列（向一个消费者发送消息）

## 延迟消息队列（向一个消费者发送消息）

## 发布订阅者模型（一次向多个消费者发送消息）