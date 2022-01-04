package main

import (
	"fmt"
	"work-place/src/github.com/garyburd/redisgo/redis"
)
func main() {

	dial, err := redis.Dial("tcp", "192.168.1.10:6379")

	defer dial.Close()
	if err!=nil {

		fmt.Println(err)
	}

	//fmt.Println(dial)

	dial.Do("select",1)
	dial.Do("SET","k1","v1")
	str,_ := dial.Do("GET","k1")
	dial.Do("lpush","k2","v2")
	fmt.Println(string(str.([]uint8)))


}
