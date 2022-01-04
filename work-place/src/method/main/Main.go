package main

import (
	"fmt"
	"strconv"
	"strings"
	"work-place/src/method"
)

/**
	go中的值传递：值传递和引用传递
		值传递：基本数据类型（int,float,string,bool,byte,数组，结构体）
		引用传递：引用数据类型（指针，slice切片，map，interface，channel管道）
 */

/**
	go中的代码执行顺序：import包中的 变量定义--> init() -->common()
 */
//在main方法之前执行，做运行前初始化作用
func init() {

	fmt.Println("init method")
}

// name := "tom" : 报错，执行语句无法在方法外执行

func main() {

	//method.Main()

	//a := closePac()
	//
	//i, s := a(1)
	//fmt.Println(i,s)
	//i2, s2 := a(2)
	//fmt.Println(i2,s2)

	//method.TestString()

	//method.TestTime()

	//method.TestNew()

	method.DealError()
}

//匿名函数第三种用法：全局匿名函数

var Func = func(i1 int,i2 int) int{
	return i1 +i2
}

func testInMethod() {

	//匿名函数第一种用法：用处较多
	res := func(i1 int,i2 int) int {

		return i1+i2
	}(10,11)

	fmt.Println(res)

	//匿名函数的第二种用法，用处较少
	res2 := func(i1 int,i2 int) int{

		return i1+i2
	}

	a := res2(10,11)
	fmt.Println(a)

	b := Func(10,11)
	fmt.Println(b)
}

func closePac() func(i1 int) (int,string) {

	var n int = 10
	var str string = ""

	return func(i1 int) (int,string) {

		i1 += n
		str += strconv.FormatInt(int64(i1),10)

		return i1,str
	}
}

func makeSuffix(suffix string) func(name string) string {

	return func(name string) string {

		if !strings.HasSuffix(name,suffix){

			return name+suffix
		}

		return name
	}
}

func testDefer() {

	//按照defer的顺序入栈，在方法执行完毕后出栈依次执行，常用于资源释放

	//入栈后会把携带的值也入栈，不会因为后续的值更改而更改
	defer fmt.Println("first")
	defer fmt.Println("second")

	res := 10+10

	fmt.Println(res)

	var i int
	for i=0;i<100;i++ {

		fmt.Println(i)
	}
}
