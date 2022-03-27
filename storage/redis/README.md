# redis 库与接口

##　github.com/go-redis/redis/v8 【推荐】


## github.com/gomodule/redigo/redis

注意：
`github.com/gomodule/redigo/redis` 如果你不传入 timeout 的值，那么默认0值的话，这两个set deadline的逻辑就跳过了。。。
如果不设置read/write timeout 会导致什么问题呢？假如网络有波动，执行一个redis 命令的时候，一直没收到服务器的响应，会导致这次请求一直没有返回，晾在那。直到redis服务器设置的超时时间到了，关闭连接，然后就会读到一个EOF的错误。
单点redis的情况，如果不设置 MaxActive，redis pool的连接数是没有上限的，问题就不会暴露出来，这对我们的服务来说，影响也不大，就是在错误日志中，会多几条redis相关的EOF日志，但是这样真的没问题么？当然有问题，如果是从redis读消息，没有设置read timeout，
一直读不到，这个协程就卡在那，迟迟不给响应，对用户来说很不好。
使用集群模式，一般redis_proxy 会限制连接数，所以redis pool 就应该用MaxActive限制池子里的最大连接数，这时候如果不设置read/write timeout，问题就来了，池子里的连接会越来越少直到没有。
因此，不管那种情况，我们都应该给redis.Dial这个方法，传入三个超时时间，DialConnectTimeout， DialReadTimeout，DialWriteTimeout。
而 `github.com/go-redis/redis/v8` 是有默认的超时时间的