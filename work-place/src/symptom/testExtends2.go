package symptom

import "fmt"

type A struct {

	name string
	age string
}

type B struct {

	name string
	age string
}

/**
	C中嵌入了AB两个结构体，且包含相同字段，调用时需要指定结构体类型，否则编译会报错
 */
type C struct {

	//you should learn to use point ,because its quicker then normal
	*A
	*B
}

func test01() {

	var c C
	c.B.name = ""

	//假设C中加入一个相同字段则就近原则调用C结构体中的字段
	d := C{
		A: &A{
			name: "jason",
			age:  "12",
		},
		B: &B{
			name: "jack",
			age:  "15",
		},
	}

	fmt.Println(*d.A,*d.B)

}