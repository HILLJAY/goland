package _interface

import "fmt"

type point interface {

	say()
}

type poIm struct {

}

//这里传入的为指针类型
func (p *poIm) say() {

	fmt.Println("say")
}