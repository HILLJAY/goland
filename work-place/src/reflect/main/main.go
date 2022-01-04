package main

import "work-place/src/reflect"

func main() {

	//var i1 int
	//i1 = 100
	//reflect.Test01(i1)

	stu := reflect.Student{
		Name: "Tom",
		Age:  16,
	}

	reflect.Test02(stu)
}
