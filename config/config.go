package config

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/26 下午10:33
 * @Desc:   定义配置
 */

// redis 配置
type RedisConfs map[string]RedisConf

type RedisConf struct {
    Network  string // 网络协议：tcp、unix
    Addr     string // 地址：127.0.0.1:6379
    Username string // 用户名
    Password string // 密码
    DB       int    // 库:0~15
    PoolSize int    // 连接池数量
}

type NacosConf struct {
    
}
