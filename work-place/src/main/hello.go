package main

import "fmt"

func main()  {

	// 1.var i int
	// 2.var i = 1
	// 3. i := 1

	// 一次申请多个变量

	/**
		有符号
		int int8 int16 int32 int64
		无符号
		uint uint8 uint16 uint32 uint64
		有符号
		rune --> 等价于int32
		无符号
		byte --> 0-255
	 */

	var f1 float32 = 1.01
	//float64精度更高，相比起float32
	var f2 float64 = 1.9000001

	var s1 int = 'a'
	fmt.Printf("%T\n",s1)
	fmt.Printf("%c\n",s1)

	fmt.Println(f2)
	fmt.Println(f1)
	//类型不匹配
	//fmt.Println(i2+i3)
	//fmt.Printf("%T",i1)

	//escapter()
}

/**
\t:制表符
\n:换行
\\: \
\r:回车
*/

func escapter() {

	var i1 int
	i1 ++

	name := "qwe"

	var i = 1

	fmt.Println(i)
	fmt.Println(name)
	fmt.Println("a"+" b")
}