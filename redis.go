package andis

/*
 * @Author: 夏华楼
 * @Date: 2020/12/31 10:12
 * @Email: Variou.xia@aishu.cn
 */
import (
	"context"
	"fmt"
	"github.com/XIAHUALOU/andis/client"
	"github.com/XIAHUALOU/andis/logger"
	r "github.com/go-redis/redis/v8"
	"runtime"
	"time"
)

//var redisClient_Once sync.Once

type RedisConf struct{}

func NewRedisCrud() *RedisConf {
	return &RedisConf{}
}

//使用默认配置
func (*RedisConf) ConfigPrepare(config interface{}) interface{} {
	var RedisConfig interface{}
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
	return RedisConfig
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

func Redis(RedisConfig interface{}) client.RedisOperator {
	var redisClient client.RedisOperator
	if v, ok := RedisConfig.(*r.Options); ok {
		redisClient = r.NewClient(v)
	} else if v, ok := RedisConfig.(*r.ClusterOptions); ok {
		redisClient = r.NewClusterClient(v)
	} else {
		redisClient = r.NewFailoverClient(RedisConfig.(*r.FailoverOptions))
	}
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Error(fmt.Sprintf(("connect error:%s"), err.Error()))
	}
	return redisClient
}
