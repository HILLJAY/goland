package reflect

import (
	"fmt"
	"reflect"
)

func TestEleForInt(b interface{}) {

	rvl := reflect.ValueOf(b)
	fmt.Println(rvl.Kind())

	/**
		rvl.Elem()类似于
		var a = 10
		var b *int = &a
		*b = 20
	*/
	rvl.Elem().SetInt(20)
	rvl.SetInt(12)
}
