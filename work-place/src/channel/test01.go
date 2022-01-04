package channel

import "fmt"

/**
管道的声明的基本使用
 */
func Test01() {

	//声明,只能存入指定类型
	var chan1 chan int
	//初始化
	chan1 = make(chan int,3)

	//var chan1 = make(chan int,3)

	fmt.Printf("chan1的值为%v,chan1的地址值为%v\n",chan1,&chan1)
	fmt.Printf("chan1的长度%v,容量%v\n",len(chan1),cap(chan1))
	//管道中加入值,管道的使用需要指定容量，否则默认为0，向管道增加数据需要在容量内，取数据需要管道中存在
	chan1<- 11

	num1 := <-chan1
	fmt.Println(num1)
}

//管道的关闭和遍历
func Test02() {

	chan1 := make(chan int,100)

	for i:=0;i<100;i++{

		chan1<- i
	}

	//关闭管道，只读无法写入管道
	close(chan1)

	for v := range chan1{

		fmt.Println(v)
	}
}