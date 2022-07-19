# elasticsearch
---

# 概述
Elasticsearch 是一个高度可扩展的开源全文本搜索和分析引擎。 它使你可以快速，近乎实时地存储，搜索和分析大量数据。 它通常用作支持具有复杂搜索功能和要求的应用程序的基础引擎技术


# 关于elasticsearch客户端包的选择

针对 Golang 的 Elasticsearch 支持，你可以访问 Elastic 的官方 github https://github.com/elastic/go-elasticsearch。

客户端主要版本与兼容的 Elasticsearch 主要版本相对应：要连接到 Elasticsearch 8.x，请使用客户端的 8.x 版本，要连接到Elasticsearch 7.x，请使用客户端的 7.x 版本

> 注意：`github.com/olivere/elastic` 第三方目前（2022-07-11）不再适用 Elasticsearch 的当前版本。`github.com/olivere/elastic` 当前仅支持到 Elasticsearch 7.x 版本


# 常见报错
1. x509: certificate signed by unknown authority

原因：https 访问elasticsearch服务器时，需要使用CA证书文件

解决：
    - 将elasticsearch服务器的`elasticsearch:/usr/share/elasticsearch/config/certs/http_ca.crt`证书文件转换为xxx.pem文件
    - xxx.pem文件的转成称字节流设置到访问elasticsearch服务器时的字段中。
        - 如：`github.com/elastic/go-elasticsearch/v8` Config结构体的`CACert`字段中
        
# 常用访问地址

## elasticsearch 
https://127.0.0.1:9200/


## kibana
http://127.0.0.1:5601/ 