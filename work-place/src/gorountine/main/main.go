package main

import (
	"fmt"
	"sync"
	"time"
	"work-place/src/gorountine"
)

func main() {

	//开启一个协程,主线程结束协程结束
	//go gorountine.Test01()
	//
	//
	//
	//for i:=0;i<10;i++{
	//
	//	fmt.Println("common i :",strconv.Itoa(i))
	//}

	for i := 0; i < 10; i++ {

		go gorountine.TestUpper(i)
	}

	time.Sleep(time.Second*10)

	lock := sync.Mutex{}
	lock.Lock()
	for i,_ := range gorountine.UpperMap{

		fmt.Println(gorountine.UpperMap[i])
	}
	lock.Unlock()
}