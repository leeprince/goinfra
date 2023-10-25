# redis 库与接口
> goredis.go【推荐】
---

# goredis.go【推荐】
> github.com/go-redis/redis/v8 
## （一）实现 `goredis_imp.go`

## 调试lua叫脚本
1. 通过 print(...) 在redis服务端输入日志
在使用 redis.Eval 执行 Lua 脚本时，由于 Redis 是在服务器端执行脚本，无法直接在客户端进行调试。但是，您可以通过在 Lua 脚本中使用 print 函数来输出变量的值，然后在 Redis 服务器的日志中查看输出结果。
```
	script := `
		local name = "John"
		local age = 30

		print("Name:", name)
		print("Age:", age)

		return "OK"
	`
	result, err := initGoredisClient().Eval(script, nil)
	if err != nil {
		fmt.Println("执行 Lua 脚本失败:", err)

	}

	fmt.Println("脚本执行结果", result)
```

2. 直接返回需要调试的变量（可以将多个变量通过Lua的表数据类型{}包含起来），并打印结果【推荐】
> 注意在 lua 脚本中只能返回一个结果，不能返回多结果，所以建议通过table的数据类型返回

> table 是 Lua 的一种数据结构用来帮助我们创建不同的数据类型，如：数组、字典等。
Lua table 使用关联型数组，你可以用任意类型的值来作数组的索引，但这个值不能是 nil。
Lua table 是不固定大小的，你可以根据自己需要进行扩容。
Lua也是通过table来解决模块（module）、包（package）和对象（Object）的。 例如string.format表示使用"format"来索引table string。

```
	script := `
		local name = "John"
		local age = 30

		print("Name:", name)
		print("Age:", age)
        
        # 
		return {name, age}
	`

	result, err := initGoredisClient().Eval(script, nil)
	if err != nil {
		fmt.Println("执行 Lua 脚本失败:", err)

	}

	fmt.Println("脚本执行结果", result)
```

# redigo.go
> github.com/gomodule/redigo/redis
> 注意：redigo_imp.go 不再继续维护实现 redis_interface.go 的接口，推荐统一使用 `github.com/go-redis/redis/v8` 库

## （一）实现 `redigo_imp.go`
注意：
`github.com/gomodule/redigo/redis` 如果你不传入 timeout 的值，那么默认0值的话，这两个set deadline的逻辑就跳过了。。。
如果不设置read/write timeout 会导致什么问题呢？
假如网络有波动，执行一个redis 命令的时候，一直没收到服务器的响应，会导致这次请求一直没有返回，晾在那。
直到redis服务器设置的超时时间到了，关闭连接，然后就会读到一个EOF的错误。
单点redis的情况，如果不设置 MaxActive，redis pool的连接数是没有上限的，
问题就不会暴露出来，这对我们的服务来说，影响也不大，就是在错误日志中，
会多几条redis相关的EOF日志，但是这样真的没问题么？当然有问题，如果是从redis读消息，没有设置read timeout，
一直读不到，这个协程就卡在那，迟迟不给响应，对用户来说很不好。
使用集群模式，一般redis_proxy 会限制连接数，所以redis pool 就应该用MaxActive限制池子里的最大连接数，
这时候如果不设置read/write timeout，问题就来了，池子里的连接会越来越少直到没有。
因此，不管那种情况，我们都应该给redis.Dial这个方法，传入三个超时时间，DialConnectTimeout， DialReadTimeout，DialWriteTimeout。
而 `github.com/go-redis/redis/v8` 是有默认的超时时间的

# Pipeline（管道）VS Lua（脚本）
## 原子性
脚本会将多个命令和操作当成一个命令在redis中执行，也就是说该脚本在执行的过程中，不会被任何其他脚本或命令打断干扰，具有原子性，在执行脚本的时候不会被其他的命令插入，因此更适合于处理事务；
而管道虽然也会将多个命令一次性传输到服务端，但在服务端执行的时候仍然是多个命令，如在执行CMD1的时候，外部另一个客户端提交了CMD9，会先执行完CMD9再执行管道中的CMD2，因此事实上管道是不具有原子性的。

## 使用场景
就场景上来说，正因为Lua脚本会被视为一个命令去执行，因为Redis是单线程执行命令的，所以我们不能在lua脚本里写过于复杂的逻辑，否则会造成阻塞，因此lua脚本适合于相对简单的事务场景；
而管道因为不具有原子性，因此管道不适合处理事务，但管道可以减少多个命令执行时的网络消耗，可以提高程序的响应速度，因此管道更适合于管道中的命令互相没有关系，不需要有事务的原子性，且需要提高程序响应速度的场景。

