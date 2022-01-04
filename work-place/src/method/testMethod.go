package method

import "fmt"

// 字母大写代表公开的public , 小写的则外包无法访问
var I int = 10

func testM01(i int,j int) (int,float64) {

	return i+j,10.0

}

func testReturning(i int)  {

	if i > 2 {

		i--
		testReturning(i)
	}

	fmt.Println(i)
}


