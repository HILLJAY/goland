package main

import (
	"time"
	"work-place/src/github.com/garyburd/redisgo/redis"
)

var pool *redis.Pool

func initPool(addr string, maxIdle int, maxActive int, idleTileOut time.Duration) {

	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {

			return redis.Dial("tcp", addr)
		},
		TestOnBorrow: nil,
		MaxIdle:      maxIdle,
		MaxActive:    maxActive,
		IdleTimeout:  idleTileOut,
		Wait:         false,
		//MaxConnLifetime: 0,
	}

}
