package cache

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"github.com/XIAHUALOU/andis/operations"
	"regexp"
	"time"
)

type CachePolicy interface {
	Before(key string) bool
	IfNil(key string, v interface{})
	SetOperation(opt *operations.Operation)
}

type CrossPolicy struct {
	keyRegex string
	Expire   time.Duration
	opt      *operations.Operation
}

//缓存穿透策略配置
func NewCrossPolicy(keyRegex string, expire time.Duration) *CrossPolicy {
	return &CrossPolicy{keyRegex: keyRegex, Expire: expire}
}

//匹配非法字符
func (self *CrossPolicy) Before(key string) bool {
	if !regexp.MustCompile(self.keyRegex).MatchString(key) {
		return false
	}
	return true
}

//设置空缓存
func (self *CrossPolicy) IfNil(key string, v interface{}) {
	self.opt.Set(key, v, operations.WithExpire(self.Expire)).Unwrap()
}

//
func (self *CrossPolicy) SetOperation(opt *operations.Operation) {
	self.opt = opt
}
