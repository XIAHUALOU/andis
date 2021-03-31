# redis-operator

@Author: 夏华楼
@Email: Variou.xia@aishu.cn

本Redis脚手架基于go-redis模块研发,语法接近rust风格,相比原生的go-redis框架使用起来更加简洁明了一些
```go
三种初始化方式
//使用默认配置只需传入redis地址即可,地址可以是
redis.NewRedisCrud().ConfigPrepare("xxx.xxx.xxx.xxx:xxxx")
//配置中心获取的多为[]byte类型
redis.NewRedisCrud().ConfigPrepare([]byte("xxx.xxx.xxx.xxx:xxxx"))
//或者使用自己的配置
conf:=redis.ConfigNew() //获取一个空白配置模板
conf.Addr = "xxx.xxx.xxx.xxx:xxxx"
redis.NewRedisCrud().ConfigPrepare(conf)
//有以下这些参数可以配置
/*Addr string
Dialer func() (net.Conn, error)
OnConnect func(*Conn) error
Password string
DB int
MaxRetries int
MinRetryBackoff time.Duration
MaxRetryBackoff time.Duration
DialTimeout time.Duration
ReadTimeout time.Duration
WriteTimeout time.Duration
PoolSize int
MinIdleConns int
MaxConnAge time.Duration
PoolTimeout time.Duration
IdleTimeout time.Duration
IdleCheckFrequency time.Duration
readOnly bool
TLSConfig *tls.Config*/
```
1.查询key
```go
//Get
operations.NewRedisOperation().Get("keyname").Unwrap() //如果报错或者未查到直接panic
operations.NewRedisOperation().Get("keyname").Unwrap_Or("报错了") //如果报错或者未查到使用Unwrap_Or的参数代替查询结果
operations.NewRedisOperation().Get("keyname").Unwrap_Or_Else(func()string{return "报错了"})//如果报错或者未查到使用自定义函数的返回值代替查询结果
//如果想要redis原生报错,使用下面的方式
search := operations.NewRedisOperation().Get("keyname")
search.Err //这样就可以获取redis原生报错
//HGet
operations.NewRedisOperation().HGet("keyname").Unwrap() //如果报错或者未查到直接panic
operations.NewRedisOperation().HGet("keyname").Unwrap_Or("报错了") //如果报错或者未查到使用Unwrap_Or的参数代替查询结果
operations.NewRedisOperation().HGet("keyname").Unwrap_Or_Else(func()string{return "报错了"})//如果报错或者未查到使用自定义函数的返回值代替查询结果
//MGet
operations.NewRedisOperation().MGet("key1","key2").Unwrap() //如果报错或者未查到直接panic
operations.NewRedisOperation().MGet("key1","key2").Unwrap_Or([]{"1","2","3"}) //如果报错或者未查到使用Unwrap_Or的参数代替查询结果参数[]interface{}
operations.NewRedisOperation().MGet("key1","key2").Unwrap_Or_Else(func()[]interface{}{return []interface{}{"1","2"}})//如果报错或者未查到使用自定义函数的返回值代替查询结果
//当使用MGet时可以使用迭代器
iter := operations.NewRedisOperation().MGet("key1","key2").Iter()
for iter.HasNext(){
	fmt.Println(iter.Next())
}
//HMGet
operations.NewRedisOperation().HMGet("key1","key2").Unwrap() //如果报错或者未查到直接panic
operations.NewRedisOperation().HMGet("key1","key2").Unwrap_Or([]string{"{}","{}","{}"}) //如果报错或者未查到使用Unwrap_Or的参数代替查询结果参数[]interface{}
operations.NewRedisOperation().HMGet("key1","key2").Unwrap_Or_Else(func()[]interface{}{return []interface{}{"1","2"}})//如果报错或者未查到使用自定义函数的返回值代替查询结果
//当使用HMGet时可以使用迭代器
iter := operations.NewRedisOperation().MGet("key1","key2").Iter()
for iter.HasNext(){
	fmt.Println(iter.Next())
}
```
2.删除key
```go
operations.NewRedisOperation().Del("key1").Unwrap() //如果报错或者未查到直接panic
operations.NewRedisOperation().Del("key1","key2").Unwrap_Or(0) //如果报错或者未成功删除全部key使用Unwrap_Or的参数代替查询结果参数，结果为删除成功的个数
operations.NewRedisOperation().Del("key1","key2").Unwrap_Or_Else(func()int64{return 1)//如果报错或者未查到使用自定义函数的返回值代替查询结果
```
3.设置key
```go
operations.NewRedisOperation().Set("name", "variou").Unwrap()//如果报错或者未查到直接panic
operations.NewRedisOperation().Set("name", "variou").Unwrap_Or(false)//存储key不成功返回false
operations.NewRedisOperation().Set("name", "variou").Unwrap_Or_Else(func() bool { return false })//存储key不成功返回func()返回值

//使用的时候可以保存operations.NewRedisOperation()再调用，直接调用其实也没有多大的影响，会影响一部分gc性能，推荐同一个方法如果多次使用的话保存一下临时对象
op := operations.NewRedisOperation()
op.Set("key1", "val1", operations.WithExpire(10*time.Second)).Unwrap()                                     //设置key1值是val1 10s过期，设置失败panic
op.Set("key2", "val2").Unwrap_Or(false)                                                                    ////设置key2值是val2 永不过期，设置失败用false代替返回
op.Set("key3", "val3", operations.WithExpire(10*time.Second)).Unwrap_Or_Else(func() bool { return false }) //设置key3值是val3 10s过期，设置失败返回func执行结果
//HSET同上套路
op.HSet("variou", "age", "18").Unwrap()
op.HSet("variou", "age", "18").Unwrap_Or(false)
op.HSet("variou", "age", "18").Unwrap_Or_Else(func() bool { return false)
//HMset
items := map[string]interface{}{}
items["age"] = "18"
items["money"] = "8000000"
items["status"] = "sleeping"
op.HMSet("variou", items, operations.WithExpire(24*time.Hour)).Unwrap() //设置过期时间24小时，设置失败panic
op.HMSet("variou", items).Unwrap_Or(false) //设置失败返回false
op.HMSet("variou", items).Unwrap_Or_Else(func() bool { return false })//设置失败返回func()
```
4.查询过期时间
```go
op.TTL("key1").Unwrap()                   //报错就panic
op.TTL("key1").Unwrap_Or(1 * time.Second) //报错就返回1s
op.TTL("key1").Unwrap_Or_Else(func() time.Duration { //报错使用func返回
	return 1 * time.Second
})
```
5.过期key
```go
op.Expire("key1", time.Second*1).Unwrap()         //设置1s过期
	op.Expire("key1", time.Second*1).Unwrap_Or(false) //设置失败返回false
	op.Expire("key1", time.Second*1).Unwrap_Or_Else(func() bool { //设置失败返回func()
		return false
	}) //设置1s过期
```
6.缓存使用
```go
//默认缓存
store := cache.NewsCache() // 使用默认缓存 json方式放回,设置缓存时间30分钟，缓存穿透策略任意字符串都可以通过，查不到就设置空值，空值过期时间20分钟
defer store.Close() //使用完放回缓存池
store.DBGetter = func() string { //设置如果缓存拿不到从DB获取数据
	return "从数据库获取的"
}
store.GetCache("key") //返回值interface{} 与Serializer_JSON这个设置有关，后续会增加新的选项，目前只支持json
//也可以修改默认cache属性，通过这个方法
func ResetDefaultCachePolicy(ser string, reg string, exp time.Duration, blank_expire time.Duration)
//获取新的缓存
cache.NewSimpleCache(operations.NewRedisOperation(), time.Second*30, Serializer_JSON, NewCrossPolicy(".*?", time.Second*20))//参数从左到右依次为redis操作符，缓存过期时间，获取缓存返回值格式，缓存穿透策略
//NewCrossPolicy方法的参数1是key的正则,匹配不到直接过滤,第二个参数是缓存穿透后设置空值得过期时间
```
