package main

import "work-place/src/interface/typeConvert"

// 接口是一种引用类型,当一个方法传入的参数为一个接口时，调用者传参为地址值
//接口体现出多态的特性
func main() {

	//_interface.Test01()
	//
	//var t1 T
	//t1 = 12
	//fmt.Println(t1)

	//interfaceTest.TestSortForStruct()

	//typeConvert.Test01()
	//phone := typeConvert.Phone{}
	camera := typeConvert.Camera{}

	typeConvert.Work(camera)
}

// 代表一个空接口，可以接受任何类型
type T interface {
	
}