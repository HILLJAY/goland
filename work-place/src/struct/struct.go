package _struct

import (
	"encoding/json"
	"fmt"
)

type Person struct {

	name string
	age int
	address string
	ptr *int
}

func testStruct() {

	var p1 Person
	p1.age = 11
	p1.address = "beijing"
	p1.name = "jason"

	var i1 int = 10
	p1.ptr = &i1


}

/**
	声明结构体的四种方式
 */
func KindOfStruct() {

	//1.
	var person Person
	person.name = "1"

	//2.
	var person2 = Person{}
	person2.name = "2"

	//3.使用结构体的引用指针，返回的是个结构体的指针，但结构体本身仍然是一个数值类型
	var person3 *Person = new(Person)
	person3.name = "3"
	//4.类似于第三种
	var person4 *Person = &Person{}
	person4.name = "4"
}

func TestPointStruct() {

	var p1 Person
	p1.name = "jason"
	p1.age = 11

	//结构体默认是值传递，如果想改为引用传递则使用指针
	var p2 *Person = &p1
	//正常来说需要使用p2的指针引用p1来实现，但goland设计者规定可以使用p2直接引用
	(*p2).age = 13

	//也是可以的
	p2.age = 14

}

type react struct {

	x int
	y int

}

type point struct {

	point1 react
	point2 react
}

/**
	结构体中的字段，在内存分配中是连续的
 */
func TestAddress() {

	var p1 point
	var p2  point
	p1.point1 = react{x: 1,y: 2}
	p2.point2 = react{x: 4,y: 5}

	fmt.Printf("x的地址%v,y的地址%v\n",&p1.point1.x,&p1.point1.y)
	fmt.Printf("point1的地址%v,point2的地址%v",&p1,&p2)
}

/**
两个结构体之间转换需要两个结构体之间的字段，个数，类型一致
 */

type Integer int

func TestTrans() {

	s1 := react{x: 1,y: 2}
	var s2 react
	//相同类型，相同字段的结构体可以转换
	s2 = s1

	//注意，go中对数据类型要求很高，如果重新声明一个数据类型进行
	var i1 Integer
	var i2 int

	//只有强转才不会报错，否则编译无法通过
	i1 = Integer(i2)

	fmt.Println(s2,i1)
}

type monster struct {

	Name string `json:"name"`
	Age int	`json:"age"`
}

func TestTag() {

	var monster1 monster
	monster1.Name = "json"
	monster1.Age = 11

	marshal, _ := json.Marshal(monster1)
	fmt.Println(marshal)
}
