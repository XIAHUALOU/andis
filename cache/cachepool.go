package cache

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"github.com/XIAHUALOU/andis/operations"
	"sync"
	"time"
)

var NewSCachePool *sync.Pool

var Serializer string
var expire time.Duration
var blank_expire_time time.Duration
var regex_string string

func init() {
	Serializer = Serializer_JSON
	expire = time.Second * 30
	regex_string = ".*?"
	blank_expire_time = time.Second * 20
	NewSCachePool = &sync.Pool{
		New: func() interface{} {
			return NewSimpleCache(operations.NewRedisOperation(), expire, Serializer, NewCrossPolicy(regex_string, blank_expire_time))
		},
	}
}

func NewsCache() *SimpleCache {
	return NewSCachePool.Get().(*SimpleCache)
}

//配置cache获取策略
func ConfigCacheNew(f func() interface{}) {
	NewSCachePool.New = f
}

func ResetDefaultCachePolicy(ser string, reg string, exp time.Duration, blank_expire time.Duration) {
	Serializer = ser
	expire = exp
	regex_string = reg
	blank_expire_time = blank_expire
}
