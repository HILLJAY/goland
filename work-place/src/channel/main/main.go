package main

import (
	"work-place/src/channel/details"
)

func main() {

	//readChan := make(chan int,100)
	//exitChan := make(chan bool,1)
	//
	//go channel.Write(readChan)
	//go channel.Read(readChan,exitChan)
	//
	//for {
	//
	//	_,ok := <-exitChan
	//
	//	if ok{
	//		break
	//	}
	//}

	details.TestRecover()
}
