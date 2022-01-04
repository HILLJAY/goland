package testMethod

import "fmt"

type Person struct {

	Name string
	Age int
}
/**
方法的传递默认也是值拷贝，方法需要指定所属类型，可以是结构体也可以是自定义的type类型
 */
func (person Person) ShowInfo() {

	person.Name = "123"
	person.Age = 12
	fmt.Printf("name=%v,age=%v",person.Name,person.Age)
}

type Circle struct {

	Radious float64
}

//需理解传入指针的内存示意图，在实际开发中传入地址的情况居多，因为效率高
func (circle *Circle) Area(radius float64) float64 {

	//如果为&circle则为circle自身的地址
	fmt.Printf("circle指针的地址%v\n",circle)

	res := 3.14 * radius * radius
	return res
}

type Student struct {

	Name string
	Age int
}

//当一个结构体重写String方法，则在输出时会调用自定义的String,类似于Java中的重写toString()
func (student *Student) String() string {

	str := fmt.Sprintf("name=%v,age=%v",student.Name,student.Age)
	return str
}
