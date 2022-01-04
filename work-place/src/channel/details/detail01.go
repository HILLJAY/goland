package details

import "fmt"

/**
管道可以声明为只读或者只写：默认情况下为全双工,如果声明只读只写，相反会报错
 */
func Test01() {

	//只写
	var chan1 chan<- int
	//只读
	var chan2 <-chan int

	chan1<- 1
	//num := <-chan1
	num := <-chan2

	fmt.Println(num)
}