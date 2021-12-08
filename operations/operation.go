package operations

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"context"
	"github.com/XIAHUALOU/andis/client"
	"github.com/XIAHUALOU/andis/outcomes"
	r "github.com/go-redis/redis/v8"
	"time"
)

type Operation struct {
	read  client.RedisOperator
	write client.RedisOperator
}

type Member = r.Z
type ZRangeBy = r.ZRangeBy

func NewRedisOperation(read client.RedisOperator, write client.RedisOperator) *Operation {
	return &Operation{read, write}
}

//string Operation
func (self *Operation) Get(key string) *outcomes.StringResult {
	return outcomes.NewStringResult(self.read.Get(context.Background(), key).Result())
}

func (self *Operation) HGet(key, field string) *outcomes.StringResult {
	return outcomes.NewStringResult(self.read.HGet(context.Background(), key, field).Result())
}

//slice Operation
func (self *Operation) MGet(keys ...string) *outcomes.SliceResult {
	return outcomes.NewSliceResult(self.read.MGet(context.Background(), keys...).Result())
}

func (self *Operation) HMGet(key string, fields ...string) *outcomes.SliceResult {
	return outcomes.NewSliceResult((self.read.HMGet(context.Background(), key, fields...).Result()))
}

//Boolean Operation
func (self *Operation) Set(key string, value interface{}, attrs ...*OperationAttr) *outcomes.BooleanResult {
	exp := OperationAttrs(attrs).Find(ATTR_EXPR).Unwrap_Or(time.Second * 0).(time.Duration) //0*time.second 代表永久
	nx := OperationAttrs(attrs).Find(ATTR_NX).Unwrap_Or(nil)
	xx := OperationAttrs(attrs).Find(ATTR_XX).Unwrap_Or(nil)

	if nx != nil {
		return outcomes.NewBooleanResult(self.write.SetNX(context.Background(), key, value, exp).Result())
	}
	if xx != nil {
		return outcomes.NewBooleanResult(self.write.SetXX(context.Background(), key, value, exp).Result())
	}
	if _, err := self.write.Set(context.Background(), key, value, exp).Result(); err != nil {
		return outcomes.NewBooleanResult(false, err)
	} else {
		return outcomes.NewBooleanResult(true, nil)
	}
}

func (self *Operation) HSet(key string, field string, value interface{}, attrs ...*OperationAttr) *outcomes.BooleanResult {
	exp := OperationAttrs(attrs).Find(ATTR_EXPR).Unwrap_Or(time.Second * 0).(time.Duration) //0*time.second 代表永久
	_, err := self.write.HSet(context.Background(), key, field, value).Result()
	if err != nil {
		return outcomes.NewBooleanResult(false, err)
	}
	if int(exp) != 0 {
		err = self.write.Expire(context.Background(), key, exp).Err()
	}
	if err != nil {
		return outcomes.NewBooleanResult(false, err)
	}
	return outcomes.NewBooleanResult(true, nil)
}

func (self *Operation) HMSet(key string, data map[string]interface{}, attrs ...*OperationAttr) *outcomes.BooleanResult {
	exp := OperationAttrs(attrs).Find(ATTR_EXPR).Unwrap_Or(time.Second * 0).(time.Duration)
	_, err := self.write.HMSet(context.Background(), key, data).Result()
	if err != nil {
		return outcomes.NewBooleanResult(false, err)
	}
	if int(exp) != 0 {
		err = self.write.Expire(context.Background(), key, exp).Err()
	}
	if err != nil {
		return outcomes.NewBooleanResult(false, err)
	}
	return outcomes.NewBooleanResult(true, nil)
}

func (self *Operation) Expire(key string, expire time.Duration) *outcomes.BooleanResult {
	return outcomes.NewBooleanResult(self.write.Expire(context.Background(), key, expire).Result())
}

//int Operation
func (self *Operation) Del(keys ...string) *outcomes.IntResult {
	return outcomes.NewIntResult(self.write.Del(context.Background(), keys...).Result())
}

//Time Operation
func (self *Operation) TTL(key string) *outcomes.TimeResult {
	return outcomes.NewTimeResult(self.read.TTL(context.Background(), key).Result())
}

/*
ZAdd
Param: key-name score member [score member…]
Explain: 将带有给定分值的成员添加到有序列表里面
*/
func (self *Operation) ZAdd(key string, members ...*Member) *outcomes.BooleanResult {
	_, err := self.write.ZAdd(context.Background(), key, members...).Result()
	if err == nil {
		return outcomes.NewBooleanResult(true, nil)
	} else {
		return outcomes.NewBooleanResult(false, err)
	}
}

/*
ZRem
Param: key-name member [member…]
Explain: 从有序集合里面移除给定的成员
*/
func (self *Operation) ZRem(key string, members ...interface{}) *outcomes.BooleanResult {
	_, err := self.write.ZRem(context.Background(), key, members...).Result()
	if err == nil {
		return outcomes.NewBooleanResult(true, nil)
	} else {
		return outcomes.NewBooleanResult(false, err)
	}
}

/*
ZCard
Param: key-name
Explain: 返回有序集合包含的成员数量
*/
func (self *Operation) ZCard(key string) *outcomes.IntResult {
	return outcomes.NewIntResult(self.read.ZCard(context.Background(), key).Result())
}

/*
ZIncrBy Operation
Param: key-name increment member
Explain: 将member成员的分值加上increment
*/
func (self *Operation) ZIncrBy(key string, increment float64, member string) *outcomes.BooleanResult {
	_, err := self.write.ZIncrBy(context.Background(), key, increment, member).Result()
	if err == nil {
		return outcomes.NewBooleanResult(true, nil)
	} else {
		return outcomes.NewBooleanResult(false, err)
	}
}

/*
ZCount
Param: key-name min max
Explain: 返回分值介于min和max之间的成员数量，包括min和max在内
*/
func (self *Operation) ZCount(key, min, max string) *outcomes.IntResult {
	return outcomes.NewIntResult(self.read.ZCount(context.Background(), key, min, max).Result())
}

/*
ZRank
Param: key-name member
Explain: 返回成员member在有序集合中的排名，成员按照分值从小到大排列
*/
func (self *Operation) ZRank(key, member string) *outcomes.IntResult {
	return outcomes.NewIntResult(self.read.ZRank(context.Background(), key, member).Result())
}

/*
ZRevRank
Param: key-name member
Explain: 返回成员member在有序集合中的排名，成员按照分值从大到小排列
*/
func (self *Operation) ZRevRank(key, member string) *outcomes.IntResult {
	return outcomes.NewIntResult(self.read.ZRevRank(context.Background(), key, member).Result())
}

/*
ZScore
Param: key-name member
Explain: 返回成员member的分值
*/
func (self *Operation) ZScore(key, member string) *outcomes.FloatResult {
	return outcomes.NewFloatResult(self.read.ZScore(context.Background(), key, member).Result())
}

/*
ZRange
Param: key-name start stop [WITHSCORES]
Explain: 返回有序集合中排名介于start和stop之间的成员，包括start和stop在内，如果给定了可选的WITHSCORES选项，那么命令会将成员的分值一并返回，成员按照分值从小到大排列
*/
func (self *Operation) ZRange(key string, start, end int64) *outcomes.StringSliceResult {
	return outcomes.NewStringSliceResult(self.read.ZRange(context.Background(), key, start, end).Result())
}

/*
ZRevRange
Param: key-name start stop [WITHSCORES]
Explain: 返回有序集合中排名介于start和stop之间的成员，包括start和stop在内，如果给定了可选的WITHSCORES选项，那么命令会将成员的分值一并返回，成员按照分值从大到小排列
*/
func (self *Operation) ZRevRange(key string, start, end int64) *outcomes.StringSliceResult {
	return outcomes.NewStringSliceResult(self.read.ZRevRange(context.Background(), key, start, end).Result())
}

/*
ZRangeByScore
Param: key-name min max [WITHSCORES] [LIMIT offset count]
Explain: 返回有序集合中分值介于min和max之间的所有成员，包括min和max在内，并按照分值从小到大的排序来返回他们
*/
func (self *Operation) ZRangeByScore(key string, by *ZRangeBy) *outcomes.StringSliceResult {
	return outcomes.NewStringSliceResult(self.read.ZRangeByScore(context.Background(), key, by).Result())
}

/*
ZRevRangeByScore
Param: key-name min max [WITHSCORES] [LIMIT offset count]
Explain: 返回有序集合中分值介于min和max之间的所有成员，包括min和max在内，并按照分值从大到小的排序来返回他们
*/
func (self *Operation) ZRevRangeByScore(key string, by *ZRangeBy) *outcomes.StringSliceResult {
	return outcomes.NewStringSliceResult(self.read.ZRevRangeByScore(context.Background(), key, by).Result())
}

/*
ZRemRangeByScore
Param: key-name key min max
Explain: 移除有序集合中分值介于min和max之间的所有成员，包括min和max在内
*/
func (self *Operation) ZRemRangeByScore(key, min, max string) *outcomes.BooleanResult {
	_, err := self.write.ZRemRangeByScore(context.Background(), key, min, max).Result()
	if err == nil {
		return outcomes.NewBooleanResult(true, nil)
	} else {
		return outcomes.NewBooleanResult(false, err)
	}
}

/*
ZRemRangeByRank
Param: key-name start stop
Explain: 移除有序集合中排名介于start和stop之间的所有成员，包括start和stop在内
*/
func (self *Operation) ZRemRangeByRank(key string, start, stop int64) *outcomes.BooleanResult {
	_, err := self.write.ZRemRangeByRank(context.Background(), key, start, stop).Result()
	if err == nil {
		return outcomes.NewBooleanResult(true, nil)
	} else {
		return outcomes.NewBooleanResult(false, err)
	}
}

// Bool Lock
func (self *Operation) Lock(key string, t time.Duration) *outcomes.BooleanResult {
	if self.TTL(key).Unwrap_Or(-1*time.Second) == -1*time.Second { //每次设置先检查锁有没有过期时间，没有添加防止死锁
		self.Expire(key, t)
	}
	return self.Set(key, "", WithNX(), WithExpire(t))
}

// Bool Release
func (self *Operation) UnLock(key string) *outcomes.BooleanResult {
	return self.Expire(key, -2*time.Second)
}
