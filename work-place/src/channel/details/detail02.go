package details

import "fmt"

func sayHello() {

	for i:=0;i<10;i++ {

		fmt.Println("Test",i)
	}
}

/**
当下面这个协程出现问题时，希望不影响其他协程的进展
 */
func TestRecover() {

	defer func() {

		if err:=recover();err!=nil {

			fmt.Println("出现错误")
		}
	}()

	//有问题
	var errMap map[int]string
	errMap[0] = "hello"

}

func TestMain() {

	go sayHello()
	go TestRecover()

	for i := 0; i < 10; i++ {

		fmt.Println("common:",i)
	}
}