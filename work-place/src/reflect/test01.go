package reflect

import (
	"fmt"
	"reflect"
)

type Student struct {

	Name string
	Age int
}

/**
操作基本数据类型
 */
func Test01(b interface{}) {

	rTyp := reflect.TypeOf(b)
	fmt.Printf("type of b: %T\n",rTyp)

	//通过ValueOf()方法获取到的变量类型为reflect.Value类型
	rVal := reflect.ValueOf(b)
	fmt.Printf("value of b:%v\n",rVal)
	//kind本质就是一个常量
	fmt.Println(rVal.Kind())

	//rVal.Int()转换时需要数据匹配
	fmt.Println(rVal.Int()+2)

	inter := rVal.Interface()
	i1 := inter.(int)
	fmt.Println(i1)

}

/**
测试结构体的反射应用
 */
func Test02(b interface{}) {

	reflect.TypeOf(b)

	//reflect.Value
	rval := reflect.ValueOf(b)
	//反射是一种运行时状态，运行时它可以知道转为空接口的类型为什么，在编译时期是不知道的
	rin := rval.Interface()
	stu := rin.(Student)

	fmt.Printf("type:%T,val:%v,name:%v",stu,stu,stu.Name)
}

func testConst() {

	//常量在声明时候，必须赋值，无法修改,常量只有基本数据类型
	const i1 = 11
}
