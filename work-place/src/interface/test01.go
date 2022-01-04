package _interface

import "fmt"
/**
	在Go语言中，接口的实现是体现在方法的实现，一个结构体实现接口的所有方法则当前结构体实现了这个接口
 */

type Usb interface {

	start()
	stop()

}

type Camera struct {

}

func (c Camera) start() {

	fmt.Println("开始")
}

func (c Camera) stop() {

	fmt.Println("停止")
}

func Test01() {

	//panic: runtime error: invalid memory address or nil pointer dereference
	//接口不能被实例
	//var i1 Usb
	//i1.start()

	var i2 Usb = Camera{}
	i2.start()
}