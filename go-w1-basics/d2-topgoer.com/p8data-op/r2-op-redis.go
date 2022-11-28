package data_op

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func R2OpRedis() {
	// part-1 conn test
	conn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		fmt.Println("fail: ", err)
		return
	}
	// defer conn.close
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn close fail")
			return
		}
	}(conn)
	// auth
	if _, err := conn.Do("auth", redisPass); err != nil {
		fmt.Println("auth fail", err)
	}

	// part-2 set/get
	key := "sgz_abc"
	_, err = conn.Do("set", key, "aaaaa")
	if err != nil {
		fmt.Println("set fail ", err)
		return
	}
	getRes, err := redis.String(conn.Do("get", key))
	if err != nil {
		fmt.Println("get fail", err)
		return
	}
	fmt.Println("get res: ", getRes)

	// part-3 expire
	// unit of sec
	if _, err := conn.Do("expire", key, 10); err != nil {
		fmt.Println("expire fail", err)
	}

	// part-4 list
	key2 := "love sida's"
	conn.Do("lpush", key2, "hair", "eye", "nose")
	lpopRes, _ := redis.String(conn.Do("lpop", key2))
	fmt.Println("lpop res:", lpopRes)
	conn.Do("expire", key2, 1)

	// part-5 hash
	key3 := "sida"
	conn.Do("hset", key3, "field-1", 11)
	conn.Do("hset", key3, "field-2", 22)
	conn.Do("expire", key3, 1)
	hgetRes, err := redis.Int(conn.Do("hget", key3, "field-1"))
	if err != nil {
		fmt.Println("hgetRes fail", err)
		return
	}
	fmt.Println("hgetRes:", hgetRes)

	// part-6 pool

}
