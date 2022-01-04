package channel

import "fmt"

func Write(readChan chan int) {

	for i:=0;i<100;i++{

		readChan<- i
		fmt.Println("write:",i)
	}

	close(readChan)
}

func Read(writeChan chan int,exitChan chan bool)  {

	for {

		v,ok := <-writeChan
		if !ok {

			break
		}
		fmt.Println("read:",v)
	}

	exitChan<- true
	close(exitChan)
}