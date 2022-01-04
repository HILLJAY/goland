package _interface

// 接口和接口之间可以extends，如果一个结构体想要实现接口则需要实现接口及接口中extends的接口方法
type Ain interface {

	test01()
}

type Bin interface {

	test02()
}

type Cin interface {

	Ain
	Bin
	test03()
}

type Cim struct {

}

func (c Cim) test03() {

}

func (c Cim) test02() {

}

func (c Cim) test01() {

}