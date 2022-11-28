package data_op

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

// main方法运行之前会先调用init方法
func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 10,
		Dial: func() (redis.Conn, error) {
			// dial
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			// auth
			if _, err := conn.Do("auth", pass); err != nil {
				fmt.Println("auth fail", err)
			}
			return conn, err
		},
	}
}

func R2OpRedisPool() {
	// 获取链接
	conn := pool.Get()
	// 归还链接
	defer conn.Close()
	conn.Do("set", "k1", 1)
	getRes, err := redis.Int(conn.Do("get", "k1"))
	if err != nil {
		fmt.Println("get fail ", err)
		return
	}
	fmt.Println("getRes:", getRes)
	conn.Do("expire", "k1", 1)

	// 关闭链接池
	pool.Close()
}
