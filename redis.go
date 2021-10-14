package andis

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"fmt"
	"github.com/XIAHUALOU/andis/logger"
	"github.com/go-redis/redis"
	"log"
	"runtime"
	"sync"
	"time"
)

var redisClient_Once sync.Once
var redisClient *redis.Client
var RedisConfig interface{}

type RedisConf struct{}

func NewRedisCrud() *RedisConf {
	return &RedisConf{}
}

//使用默认配置
func (*RedisConf) ConfigPrepare(config interface{}) {
	switch config.(type) {
	case string:
		RedisConfig = &redis.Options{
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
		RedisConfig = &redis.Options{
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
	case *redis.Options:
		RedisConfig = config.(*redis.Options)
	case *redis.FailoverOptions:
		RedisConfig = config.(*redis.FailoverOptions)
	}
}

//ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
//user_credential_login_count := &userCredentialLoginCount{
//	ID:             "hast_users",
//	AcountPwdCount: 1,
//	TokenCount:     2,
//}
//返回一个全新的配置
func ConfigNew() *redis.Options {
	return &redis.Options{}
}

func Redis() *redis.Client {
	redisClient_Once.Do(func() {
		if v, ok := RedisConfig.(*redis.Options); ok {
			redisClient = redis.NewClient(v)
		} else {
			redisClient = redis.NewFailoverClient(RedisConfig.(*redis.FailoverOptions))
		}
		pong, err := redisClient.Ping().Result()
		if err != nil {
			logger.Error(fmt.Sprintf(("connect error:%s"), err.Error()))
		}
		log.Println(pong)
	})
	return redisClient
}
