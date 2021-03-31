package cache

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"encoding/json"
	"github.com/XIAHUALOU/andis/operations"
	"time"
)

const (
	Serializer_JSON = "Json"
)

type DBGetFunc func() string

type SimpleCache struct {
	Operation  *operations.Operation
	Expire     time.Duration
	DBGetter   DBGetFunc
	Serializer string
	Policy     CachePolicy
}

func NewSimpleCache(operation *operations.Operation, expire time.Duration, serializer string, policy CachePolicy) *SimpleCache {
	policy.SetOperation(operation)
	return &SimpleCache{Operation: operation, Expire: expire, Serializer: serializer, Policy: policy}
}

//设置缓存
func (self *SimpleCache) SetCache(key string, value interface{}) {
	self.Operation.Set(key, value, operations.WithExpire(self.Expire)).Unwrap()
}

func (self *SimpleCache) GetCache(key string) (ret interface{}) {
	if self.Policy != nil {
		if !self.Policy.Before(key) {
			return
		}
	}
	ret = self.Operation.Get(key).Unwrap_Or_Else(self.DBGetter) //redis和Db都没有拿到设置空缓存
	if ret.(string) == "" && self.Policy != nil {
		self.Policy.IfNil(key, "")
	} else {
		self.SetCache(key, ret)
	}
	if self.Serializer == Serializer_JSON {
		ret, _ = json.Marshal(ret)
	}
	return
}

func (self *SimpleCache) Close() {
	NewSCachePool.Put(self)
}
