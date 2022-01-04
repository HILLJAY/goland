package typeConvert

import "fmt"

type Usb interface {

	start()
	end()
}

type Camera struct {

}

type Phone struct {

}

func (c Camera) start()  {

	fmt.Println("camera 开始工作")
}

func (c Camera) end() {

	fmt.Println("camera 结束工作")
}

func (p Phone) start()  {

	fmt.Println("Phone 开始工作")
}

func (p Phone) end() {

	fmt.Println("Phone 结束工作")
}

func (p Phone) call() {

	fmt.Println("Phone 可以打电话")
}

func Work(usb Usb) {

	usb.start()
	usb.end()

	//如果不加判断条件，在进行类型断言之后如果传入的类型不匹配会报错（panic），所以需要显示赋值一个判断条件
	phone,ok := usb.(Phone)

	if ok {
		phone.call()
	}else {

		fmt.Println("继续工作")
	}
}

//循环入参根据不同类型显示不同的结果，这个是类型断言的最佳实践，本质就是确认接口类型
func choseType(items ...interface{}) {

	for i,v := range items{

		switch v.(type) {

		case bool:
			fmt.Printf("parame #%d is a value of bool,value is %v",i,v)
		case int:
			fmt.Printf("parame #%d is a value of int,value is %v",i,v)
		default:
			fmt.Printf("parame #%d's type is unknown")
		}
	}
}