package main

import (
	"fmt"
	"strconv"
)

func main() {

	//sre3 := `
	//
	//	this is just try the sign
	//
	//	`
	//fmt.Println(sre3)
	//
	////+做字符串拼接需要放在上一行的末尾，不然会报错
	//str4 := "123"+"456"+
	//		"123"
	//
	//fmt.Println(str4)

	//var i1 int = 12222
	//var f1 float32 = float32(i1)
	//var f2 float64 = float64(i1)
	//var u1 uint = uint(i1)
	//
	//var i2 int8 = int8(i1)
	//
	//fmt.Printf("i1=%v,f1=%v,f2=%v,u1=%v\n",i1,f1,f2,u1)
	//
	////数据类型转换可以大-->小，也可以小-->大，大到小编译时不会报错，运行结果会以溢出的形式返回
	//fmt.Println(i2)
	////转换数据类型只是把数值转换过去，而非改变原来的数据类型
	//fmt.Printf("%T",i1)

	testTrans()

	testTransStr()
	fmt.Println()
	strToBasic()
}

func testTrans() {

	var n1 int32 = 12
	var n2 int8
	var n3 int8

	n2 = int8(n1)+127
	//因为int8的范围为-128 - 128 加上的数字已经超过128了所以编译报错
	//n3 = int8(n1)+128

	fmt.Println(n2,n3)
}

/**
	基本数据类型转换为String
	String转换为基本数据类型
 */
func testTransStr() {

	 var i1 int32 = 99
	 var f1 float64 = 99.1
	 var b bool = true
	 var b2 byte = '1'
	 var str string

	 str = fmt.Sprintf("%d",i1)
	 str = fmt.Sprintf("%f",f1)
	 str = fmt.Sprintf("%b",b)
	 str = fmt.Sprintf("%c",b2)

	 fmt.Println(str)
	 fmt.Printf("%T,%q",str,str)
}

func strToBasic() {

	var s1 string = "123"
	var n3 int64
	n3, _ = strconv.ParseInt(s1,10,64)

	fmt.Printf("%T\n",n3)
	fmt.Println(n3)

	float, _ := strconv.ParseFloat(s1, 64)
	fmt.Println(float)
	fmt.Printf("%T\n",float)

	var str string = "hello"
	var i1 int64
	var f1 float64
	var b1 bool

	//如果转换不成功。则会定义为默认值
	f1,_ = strconv.ParseFloat(str,64)
	fmt.Println(f1)
	i1,_ = strconv.ParseInt(str,10,32)
	fmt.Println(i1)
	b1,_ = strconv.ParseBool(str)
	fmt.Println(b1)
}