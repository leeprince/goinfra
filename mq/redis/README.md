# redis 实现消息队列

> 文件结构
```
├── delay_mq_interface.go // 延迟队列接口
├── delay_mq_zadd_zrangebyscore_zrem_imp.go // zadd + zrangebyscore&zrem 实现延迟队列接口
├── delay_mq_zadd_zrangebyscore_zrem_imp_test.go
├── mq_interface.go // 普通消息队列接口
├── mq_push_pop_imp.go // lpush + rpop (默认，左近右出)|| rpush + lpop 实现普通消息队列接口 
└── mq_push_pop_imp_test.go
```

## 普通消息队列
> 多个发布者，一个消费者

## 延迟消息队列
> 多个发布者，一个消费者

## 发布订阅者模型
> 多个发布者，多个消费者订阅