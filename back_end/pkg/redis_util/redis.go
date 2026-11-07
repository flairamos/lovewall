package redisutil

import (
	"errors"
	"fmt"
	"shoollove/config"
	"sync"

	"github.com/gomodule/redigo/redis"
)

var lock sync.Mutex

var Client *CustomRedisClient

type CustomRedisClient struct {
	redis.Conn
}

func Init() {
	conf := config.Config.Redis
	conn, err := redis.Dial("tcp", conf.Host)
	if err != nil {
		panic(errors.New("conn redis failed"))
	}
	_, err = conn.Do("auth", conf.Password)
	if err != nil {
		panic(errors.New("redis auth failed"))
	}

	// 选择第一个数据库
	_, err = conn.Do("SELECT", conf.Database)
	if err != nil {
		panic(errors.New("redis select database failed"))
	}
	Client = &CustomRedisClient{
		conn,
	}
	fmt.Println("conn redis success")
}

func (c *CustomRedisClient) Cmd(commandName string, args ...interface{}) (reply interface{}, err error) {
	lock.Lock()
	defer lock.Unlock()
	return c.Do(commandName, args...)
}
