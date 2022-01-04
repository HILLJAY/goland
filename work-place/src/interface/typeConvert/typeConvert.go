package typeConvert

import "fmt"

type empty interface {

}

func Test01() {

	var emp empty
	var i1  = 10
	emp = i1

	i2 := 11
	//测试类型断言,接口类型无法直接赋值给一个具体的类型，需要类型断言来进行转换
	i2 = emp.(int)
	fmt.Printf("i2=%v",i2)
}