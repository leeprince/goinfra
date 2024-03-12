# 对接 WordPress REST API 
---

# 功能
1、上传文章（帖子）

# 关于 wordpress REST API

## 注意
在 `设置` - `固定链接` 中一定不能使用 **朴素的** 的固定链接结构

## 插件
要启用 Basic Auth 相关插件：`WordPress REST API Authentication`

## curl 测试
```
# --request的参数区分大小写：只能大写
curl --request GET --user "用户名:用户名密码" https://itgogogo.cn/wp-json/wp/v2/users

# 发现不需要账号和密码也是可以获取到用户信息的
curl --request GET https://itgogogo.cn/wp-json/wp/v2/users
```