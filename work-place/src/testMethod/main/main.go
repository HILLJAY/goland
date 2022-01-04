package main

import (
	"fmt"
	"work-place/src/testMethod"
)

func main() {

	//var p testMethod.Person
	//p.Name = "tom"
	//p.Age = 123
	//p.ShowInfo()
	//
	//fmt.Printf("name=%v,age=%v",p.Name,p.Age)

	//var c1 testMethod.Circle
	//c1.Radious = 10
	////因为传入的是一个地址，所以按理来说因该(&c1).Area((&c1).Radious),只不过底层做了优化
	//res := c1.Area(c1.Radious)
	//
	//fmt.Println(res)
	//fmt.Printf("c1的地址%v",&c1)

	//var str testMethod.Student
	//str.Name = "gyh"
	//str.Age = 11
	//
	//fmt.Println(&str)

	//在创建结构体的时候，可以直接指定字段值，这里返回指针类型
	//
	//var hero1 = &testMethod.Hero{Name: "jason",Age: 12}
	//hero2 := &testMethod.Hero{Name: "jack",Age: 13}
	//
	//hero1.ShowInfo()
	//hero2.ShowInfo()
	//
	////hero1:&{jason 12},hero2:&{jack 13}如果直接输出则带有&符号
	//fmt.Printf("hero1:%v,hero2:%v",*hero1,*hero2)

	student := testMethod.StuBeanFactory("gyh", 12)

	student.SetName("json")
	//println(student.GetName())
	//返回值为一个指针类型，输出结果为所存如的地址值
	fmt.Printf("返回值类型%T,返回值结果%v",student,*student)

}


