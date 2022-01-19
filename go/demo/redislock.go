package redislock

import (
	"common/settings"
	"fmt"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type RedisLock struct {
	rs    *redsync.Redsync
	mutex *redsync.Mutex
}

//Cluster
type RedisInfo struct {
	Addr         string
	Password     string
	PoolSize     int
	MinIdleConns int
}

// Redis锁-集群模式（Redis个数必须是奇数）
func NewRedisLockByCluster(redisList []*settings.RedisConfig) *RedisLock {
	count := len(redisList)
	if count%2 == 0 { //Redis个数必须是奇数
		return nil
	}

	clientPools := make([]redis.Pool, count)
	for idx := 0; idx < count; idx++ {
		redis_addr := fmt.Sprintf("%s:%d", redisList[idx].Host, redisList[idx].Port)
		client := goredislib.NewClient(&goredislib.Options{
			Addr:         redis_addr, //"localhost:6379"
			Password:     redisList[idx].Password,
			PoolSize:     redisList[idx].PoolSize,
			MinIdleConns: redisList[idx].MinIdleConns,
		})

		pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
		clientPools[idx] = pool
	}

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(clientPools...)

	return &RedisLock{
		rs: rs,
	}

}

// Redis锁-单机模式
func NewRedisLock(addr string, password string, poolSize int, minIdleConns int) *RedisLock {

	client := goredislib.NewClient(&goredislib.Options{
		Addr:         addr, //"localhost:6379"
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
	})

	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	return &RedisLock{
		rs: rs,
	}
}

// 申请Redis锁
func (rl *RedisLock) Acquire(mutexname string, expiry time.Duration) error {

	var options []redsync.Option
	options = append(options, redsync.WithExpiry(expiry))
	options = append(options, redsync.WithTries(3))
	options = append(options, redsync.WithRetryDelay(time.Microsecond*50))
	options = append(options, redsync.WithDriftFactor(0.8))
	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	//mutexname := "my-global-mutex"
	rl.mutex = rl.rs.NewMutex(mutexname, options...)
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := rl.mutex.Lock(); err != nil {
		return err
	}
	return nil
}

// 释放Redis锁
func (rl *RedisLock) Release() error {

	if _, err := rl.mutex.Unlock(); err != nil {
		return err
	}
	return nil
}
