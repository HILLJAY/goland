package method

import "fmt"

var k int = 10

func testArr(i *int) int{

	*i = 100

	return *i
}

func getSum(s1 int,s2 int) int {

	return s1+s2
}

//支持自定义类型，实质上其实是给一个类型起别名
type myType func(int,int) int

func toGetSum(tp myType,num1 int,num2 int) int {

	return tp(num1,num2)
}

//支持对返回值进行命名
func returnTest(num1 int,num2 int) (sum int , sub int)  {

	sum = num1+num2
	sub = num2-num1

	return sum,sub
}

func mutNums(num1 int, vars ...int) {

	sum := num1

	for i := 0; i < len(vars); i++ {

		sum += vars[i]
	}

	fmt.Println(sum)
}

func swap(num1 *int, num2 *int) {

	temp := *num1
	*num1 = *num2
	*num2 = temp

	fmt.Printf("%T\n",temp)
}

func Main() {

	//可以用一个变量接受一个方法类型，方法也是一种数据类型，同时可以作为形参
	a := mutNums

	fmt.Printf("%T",a)

	mutNums(10,0,90,10)

	num1 := 10
	num2 := 20
	swap(&num1,&num2)

	fmt.Println("a=",num1,"b=",num2)
}

