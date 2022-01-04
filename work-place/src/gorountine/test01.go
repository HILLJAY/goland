package gorountine

import (
	"fmt"
	"runtime"
	"strconv"
)

func Test01() {

	i := 100
	for ;i>=0;i--{

		fmt.Println("test i :",strconv.Itoa(i))
	}
}

func Test02() {

	num := runtime.NumCPU()
	fmt.Println(num)

	runtime.GOMAXPROCS(7)
}