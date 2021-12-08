package client

import (
	"context"
	r "github.com/go-redis/redis/v8"
	"time"
)

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
