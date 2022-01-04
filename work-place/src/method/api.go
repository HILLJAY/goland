package method

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func TestString() {

	var str string = "123中国"

	//获取每一个字符对应的ASCC码，需要重新编码形成字符
	r := []rune(str)
	for i := 0;i<len(r);i++ {

		fmt.Printf("%c",r[i])
	}

	//字符串转整数，整数转字符串
	atoi, _ := strconv.Atoi("123")
	fmt.Println(atoi)

	itoa := strconv.Itoa(123)
	fmt.Println(itoa)

	//string --> bytes
	byte := []byte("hello")
	fmt.Printf("bytes=%v",byte)

	//bytes --> string
	s := string(byte)
	fmt.Printf("\n%v\n",s)

}

func test03(i int) {

	var str string
	for i := i;i<10000;i++ {

		str += strconv.Itoa(i)
	}
}

func TestTime() {

	//now := time.Now()
	//fmt.Printf("time is %v the type is %T",now,now)

	start := time.Now().Unix()
	test03(0)
	end := time.Now().Unix()

	fmt.Println(end - start)
}

func TestNew() {

	num2 := new(int)

	*num2 = 100
	fmt.Printf("num2的值%v,num2的地址%v,num2的类型%T,num2指向的值%v",num2,&num2,num2,*num2)
}

func TestError() {

	var num1 int = 10
	var num2 int = 0

	defer func() {

		err := recover()

		if err!=nil{

			fmt.Println(err)
		}
	}()

	res := num1/num2

	fmt.Println(res)

}

func TestYoursError(filename string) error {

	if filename == ""{

		return errors.New("文件为空")
	}

	return nil
}

func DealError() {

	error := TestYoursError("")

	if error!=nil{

		panic(error)
	}else {

		fmt.Println("继续运行")
	}

}