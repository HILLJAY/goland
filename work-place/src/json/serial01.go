package json

import (
	"encoding/json"
	"fmt"
)

type Monster struct {

	//通过反射机制使结构体序列化后的标签为指定的标签
	Name string `json:"name"`
	Age int	`json:"age"`
	Address string `json:"adddress"`
}

/**
	测试将结构体,map,切片序列化为一个json数据,基本数据类型转换为json后意义不大，不是key-value的形式
 */
func Test01() {

	monster := Monster{
		Name : "year",
		Age : 100,
		Address : "spring",
	}

	//空接口可以传入所有类型的数据，传入地址值代表一个指针类型
	byte,err := json.Marshal(&monster)

	if err!=nil {

		fmt.Println("error:",err)
	}

	fmt.Printf("序列化后的结果%v",string(byte))
}

func Test02() {

	var jsonMap = make(map[string]interface{})
	jsonMap["name"] = "jack"
	jsonMap["age"] = 12
	jsonMap["address"] = "nanjing"

	byte,err := json.Marshal(jsonMap)

	if err!=nil {

		fmt.Println("err:",err)
	}

	fmt.Printf("%v",string(byte))

}

func Test03() {

	var slice []map[string]interface{}
	map1 := make(map[string]interface{})
	map1["name"] = "jack"
	map1["age"] = 15
	map1["address"] = "北京"

	map2 := make(map[string]interface{})
	map2["name"] = "tom"
	map2["age"] = 17
	map2["address"] = "南京"

	slice = append(slice,map1,map2)

	byte,err := json.Marshal(slice)

	if err!=nil {
		fmt.Println(err)
	}

	fmt.Printf("%v",string(byte))
}