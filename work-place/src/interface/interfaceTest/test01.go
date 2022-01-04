package interfaceTest

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {

	name string
	age int
}

//声明一个结构体切片，然后实现sort接口里面的方法
type heroSlice []Hero

func (h heroSlice) Len() int {

	return len(h)
}

func (h heroSlice) Swap(i, j int) {

	temp := h[i]
	h[i] = h[j]
	h[j] = temp
}

func (h heroSlice) Less(i, j int) bool {

	return h[i].age<h[j].age
}

func TestSortForStruct() {

	var slice heroSlice = make([]Hero,2)
	for i:=0;i<10;i++{

		var hero  = Hero{
			name: fmt.Sprintf("英雄编号:%v",rand.Intn(100)),
			age:  rand.Intn(100),
		}

		//往切片中添加hero
		slice = append(slice,hero)
	}

	fmt.Printf("排序之前%v\n",slice)

	sort.Sort(slice)
	fmt.Printf("排序之后%v\n",slice)

}