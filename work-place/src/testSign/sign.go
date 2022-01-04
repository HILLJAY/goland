package testSign

import "fmt"

func test01() {

	i := 10

	fmt.Println(i)
}

func Test02() {

	//// go语言中的/和%：
	////	对于/ ： 如果是整数的话则结果只取整数部分，如果需要为浮点型则为12.0
	////	对于%
	//var i int = 10
	//var b float64 = 12.0
	//
	//i = i%5
	//b = b%6
	//
	////++ -- 只能单独使用 只有i++ , b++
	//i++
	//b++
	//
	//fmt.Println(i)

	var days int = 97
	var week int = days/7
	var day int = days%7

	fmt.Printf("%d个星期零%d天",week,day)

}

func GetTem(huashi float64) {

	var tem float64 = 5.0/9*(huashi-100)

	fmt.Println(tem)
	fmt.Printf("%v对应的摄氏温度为%v",huashi,tem)

}

func SwapNum() {

	var a int = 10
	var b int = 20

	//a = 30
	a = a+b
	//b = 10
	b = a-b
	//a = 20
	a = a-b

	fmt.Printf("a=%v b=%v",a,b)

	var ptr *int = &a
	fmt.Printf("%v的值为%v",ptr,*ptr)
}

func TestScan() {

	var name string
	var age int
	var address string
	var isPass bool

	fmt.Println("请输入名字")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入地址")
	fmt.Scanln(&address)
	fmt.Println("判断是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("%v - %v - %v - %v",name,age,address,isPass)

}

func TestSwitch() {

	var c string
	fmt.Println("请输入一个字符")
	//fmt.Scanf("%c",&c)

	var i int =0
	i,_=fmt.Scanln(&c)

	switch c {

	case "a":
		fmt.Println("A")
	}

	fmt.Println(i)
}

func testFor() {

	for i := 10; i >= 0; i++ {

		fmt.Println("hello",i)
	}
	
}