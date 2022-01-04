package main

import "work-place/src/github.com/garyburd/redisgo/redis"

func main() {

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {

			return redis.Dial("tcp","192.168.1.10:6379")
		},
		//TestOnBorrow:    nil,
		MaxIdle:         8,
		MaxActive:       1,
		IdleTimeout:     10,
		Wait:            false,
		//MaxConnLifetime: 0,
	}

	con := pool.Get()
	defer con.Close()

	con.Do("set","k3","v3")

	//如果连接池关闭则无法取到连接
	//pool.Close()
}
