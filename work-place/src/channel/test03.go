package channel

import "fmt"

func MainTest() {

	writeChan := make(chan int,100)
	priNum := make(chan int,50)
	exitChan := make(chan string,4)

	go WritePrime(writeChan, 100)
	for i := 0; i < 4; i++ {

		go PrimeNum(exitChan, priNum, writeChan)
	}

	for  {

		if len(exitChan) == 4{

			close(priNum)
			break
		}
	}

	for {

		v,ok := <-priNum

		if !ok{
			break
		}

		fmt.Printf("素数为:%v\n",v)
	}
}
func WritePrime(writeChan chan int,num int){

	for i:=0;i<num;i++ {

		writeChan<- i
	}

	close(writeChan)
}

func PrimeNum(exit chan string,priNum chan int,writeChan chan int) {

	for {

		v,ok := <-writeChan

		if !ok {

			exit<- "ok"
			break
		}

		if JudgeNum(v) {

			priNum<- v
		}
	}

}

func JudgeNum(num int) bool {

	if num<2{

		return false
	}

	for i := 2; i < num; i++ {

		if num%i == 0{

			return false
		}
	}

	return true
}