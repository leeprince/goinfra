# 日志


## 使用方式


### 默认
默认直接调用

### 自定义配置
**支持自定义的方法：logger_option.go**

1. 更快的json序列化方式 `SetFormatterJsonInter`
2. 保存日志记录到文件 `SetOutputFile`
3. 添加钩子实现 logger 的记录调用日志者的标记 `AddHookOfReportCaller`
> 原始的 SetReportCaller，需特殊使用方可记录实际调用日志者的标记。具体说明：logger.go@SetReportCaller
4. // TODO: 日志切分 - prince@todo 2022/3/15 下午10:08
5. // TODO: logID - prince@todo 2022/3/15 下午10:09
6. // TODO: sentry - prince@todo 2022/3/15 下午10:09

