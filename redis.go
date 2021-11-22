package andis

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"context"
	"fmt"
	"github.com/XIAHUALOU/andis/logger"
	r "github.com/go-redis/redis/v8"
	"log"
	"runtime"
	"sync"
	"time"
)

var redisClient_Once sync.Once
var redisClient RedisOperator
var RedisConfig interface{}

type RedisConf struct{}

func NewRedisCrud() *RedisConf {
	return &RedisConf{}
}

//使用默认配置
func (*RedisConf) ConfigPrepare(config interface{}) {
	switch config.(type) {
	case string:
		RedisConfig = &r.Options{
			Network:            "tcp",
			Addr:               config.(string),
			DialTimeout:        5 * time.Second,
			ReadTimeout:        3 * time.Second,
			WriteTimeout:       3 * time.Second,
			PoolSize:           4 * runtime.NumCPU(),
			PoolTimeout:        4 * time.Second,
			IdleTimeout:        5 * time.Second,
			IdleCheckFrequency: 60 * time.Second,
		}
	case []byte:
		RedisConfig = &r.Options{
			Network:            "tcp",
			Addr:               string(config.([]byte)),
			DialTimeout:        5 * time.Second,
			ReadTimeout:        3 * time.Second,
			WriteTimeout:       3 * time.Second,
			PoolSize:           4 * runtime.NumCPU(),
			PoolTimeout:        4 * time.Second,
			IdleTimeout:        5 * time.Second,
			IdleCheckFrequency: 60 * time.Second,
		}
	case *r.Options:
		RedisConfig = config.(*r.Options)
	case *r.FailoverOptions:
		RedisConfig = config.(*r.FailoverOptions)
	case *r.ClusterClient:
		RedisConfig = config.(*r.ClusterOptions)
	}
}

//ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
//user_credential_login_count := &userCredentialLoginCount{
//	ID:             "hast_users",
//	AcountPwdCount: 1,
//	TokenCount:     2,
//}
//返回一个全新的配置
func ConfigNew() *r.Options {
	return &r.Options{}
}

type RedisOperator interface {
	Get(context.Context, string) *r.StringCmd
	HGet(context.Context, string, string) *r.StringCmd
	MGet(context.Context, ...string) *r.SliceCmd
	HMGet(context.Context, string, ...string) *r.SliceCmd
	Set(context.Context, string, interface{}, time.Duration) *r.StatusCmd
	SetNX(context.Context, string, interface{}, time.Duration) *r.BoolCmd
	SetXX(context.Context, string, interface{}, time.Duration) *r.BoolCmd
	HSet(context.Context, string, ...interface{}) *r.IntCmd
	HMSet(context.Context, string, ...interface{}) *r.BoolCmd
	Expire(context.Context, string, time.Duration) *r.BoolCmd
	Del(context.Context, ...string) *r.IntCmd
	TTL(context.Context, string) *r.DurationCmd
	ZAdd(context.Context, string, ...*r.Z) *r.IntCmd
	ZRem(context.Context, string, ...interface{}) *r.IntCmd
	ZCard(context.Context, string) *r.IntCmd
	ZIncrBy(context.Context, string, float64, string) *r.FloatCmd
	ZCount(context.Context, string, string, string) *r.IntCmd
	ZRank(context.Context, string, string) *r.IntCmd
	ZRevRank(context.Context, string, string) *r.IntCmd
	ZScore(context.Context, string, string) *r.FloatCmd
	ZRange(context.Context, string, int64, int64) *r.StringSliceCmd
	ZRevRange(context.Context, string, int64, int64) *r.StringSliceCmd
	ZRangeByScore(context.Context, string, *r.ZRangeBy) *r.StringSliceCmd
	ZRevRangeByScore(context.Context, string, *r.ZRangeBy) *r.StringSliceCmd
	ZRemRangeByScore(context.Context, string, string, string) *r.IntCmd
	ZRemRangeByRank(context.Context, string, int64, int64) *r.IntCmd
	Ping(ctx context.Context) *r.StatusCmd
}

func Redis() RedisOperator {
	redisClient_Once.Do(func() {
		if v, ok := RedisConfig.(*r.Options); ok {
			redisClient = r.NewClient(v)
		} else if v, ok := RedisConfig.(*r.ClusterOptions); ok {
			redisClient = r.NewClusterClient(v)
		} else {
			redisClient = r.NewFailoverClient(RedisConfig.(*r.FailoverOptions))
		}
		pong, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			logger.Error(fmt.Sprintf(("connect error:%s"), err.Error()))
		}
		log.Println(pong)
	})
	return redisClient
}
