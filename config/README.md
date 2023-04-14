# 分布式配置中心
---

# 一、Nacos
[nacos文档](https://nacos.io/zh-cn/docs/what-is-nacos.html)
## （一）概念
Nacos 致力于帮助您发现、配置和管理微服务。Nacos 提供了一组简单易用的特性集，帮助您快速实现动态服务发现、服务配置、服务元数据及流量管理。
Nacos 帮助您更敏捷和容易地构建、交付和管理微服务平台。 Nacos 是构建以“服务”为中心的现代应用架构 (例如微服务范式、云原生范式) 的服务基础设施。

## （二）`动态配置服务`实现分布式配置中心
动态配置服务可以让您以中心化、外部化和动态化的方式管理所有环境的应用配置和服务配置。
动态配置消除了配置变更时重新部署应用和服务的需要，让配置管理变得更加高效和敏捷。
配置中心化管理让实现无状态服务变得更简单，让服务按需弹性扩展变得更容易。

Nacos 提供了一个简洁易用的UI (控制台样例 Demo) 帮助您管理所有的服务和应用的配置。
Nacos 还提供包括配置版本跟踪、灰度发布、一键回滚配置以及客户端配置更新状态跟踪在内的一系列开箱即用的配置管理特性，帮助您更安全地在生产环境中管理配置变更和降低配置变更带来的风险。


## （三）部署
- docker 部署：https://nacos.io/zh-cn/docs/quick-start-docker.html
    - docker hub：https://hub.docker.com/r/nacos/nacos-server
    - 源码地址：https://github.com/nacos-group/nacos-docker.git
    - 开发/测试环境：MODE=standalone
        其他参数无需再配置
    - 默认账号密码。账号：nacos 密码：nacos    
    
    
# xconf 分布式配置中心
> 参考：https://github.com/stack-labs/XConf


# 变更记录
1. `NewNacosClient` 必填参数与可选参数分开
2. 新增 `MustNewNacosClient`，出现错误抛出异常
