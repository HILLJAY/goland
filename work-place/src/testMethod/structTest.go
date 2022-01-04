package testMethod

import "fmt"

type Hero struct {

	Name string
	Age int
}

func (hero *Hero) ShowInfo() {

	fmt.Printf("name=%v,age=%v\n",hero.Name,hero.Age)
}

//

/**
测试工厂模式：当一个结构体的首字母小写的时候，我们可以通过工厂模式来返回一个实例
 */

type student struct {

	name string
	age int
}

func (s *student) GetName() string {

	return s.name
}

func (s *student) SetName(name string)  {

	s.name = name
}

//这里工厂模式的实现需要使用函数，不能使用方法，因为方法的调用需要结构体，而此处需求为当结构体小写的时候如何获取结构体实例
func StuBeanFactory(name string,age int) *student {

	//temp为一个指针类型
	var temp = &student{
		name: name,
		age:  age,
	}

	return temp
}