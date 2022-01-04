package json

import (
	"encoding/json"
	"fmt"
)

type Stu struct {

	Name string
	Age int
	Address string
}

func TestUnSer() {

	monster := Stu{
		Name:    "jack",
		Age:     12,
		Address: "南京",
	}

	byte,err := json.Marshal(monster)
	if err != nil {

		fmt.Println(err)
	}

	var monster02 Stu
	er := json.Unmarshal(byte,&monster02)
	if er!=nil{

		fmt.Println(er)
	}

	fmt.Println(monster02)
}