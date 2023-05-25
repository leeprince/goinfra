# go 基础服务
---

# 一、日志
> plog


# 二、存储
> storage
## （一）Mysql
## （二）Reids
## （二）Elasticsearch // TODO:  - prince@todo 2022/6/6 下午10:30

# 三、分布式锁
> lock
## （一）redis 实现

# 四、HTTP 服务
> http
## （一）http 客户端 // TODO: 升级后的改造 - prince@todo 2023/4/8 22:57

# 五、告警中心
> alert
## sentry

# 六、动态配置管理
> manage_config
## Nacos

# 七、服务管理
> manage_service
## （一）Etcd
## （二）Nacos

# 八、常量定义
> consts
## （一）环境
## （二）Mime

# 九、消息队列
> mq
## （一）Redis
## （二）RabbimMQ

# 十、分布式链路追踪
> trace
## （一）opentracing


# 十一、任务
> task
##（一）重试任务

##（二）定时任务

##（三）推送任务

##（四）并发任务
### 1. 并发执行，并发执行过程中，遇到错误不终止所有并发任务，待所有并发任务结束后判断是否存在错误，并对错误进行处理
### 2. 并发执行，并发执行过程中，遇到错误终止所有并发任务

# 十二、websocket // TODO:  - prince@todo 2022/6/6 上午1:13
> websocket

# 十三、资源
> resource

# 十四、工具类（业务无关）
> utils
>   - `utils` ：通用的且与项目业务无关的类的组合；可供其他项目使用。如：字符串工具类,文件工具类等。`tools` ：当前项目中通用的业务类的组合；仅能当前项目使用。如：用户校验工具类,支付工具类等
## （一）CICD
> 持续集成 CI(Continuous Integration)
> 持续交付 CD (Continuous Delivery)
## （二）jsonpb
## （三）proto
## （四）tablestruct
## （五）切片
## （六）时间
## （七）uuid（唯一ID）
## （八）字符串
### 1. 驼峰命名法 & 蛇形命名法


# 十五、安全
> security

# 十六、状态码管理
> code

# 十七、键值对管理
> kv

# 十八、测试工具

# 十九、Mock管理

# 二十、monitor

# 二一、rpa


# todo priority
# ES
