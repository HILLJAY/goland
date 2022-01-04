package arr

import (
	"fmt"
	"sort"
)

func TestArr01() {

	var arr [3]int
	arr[0] = 1
	arr[1] = 2

	fmt.Printf("arr[0]的地址%v,arr[1]的地址%v,arr[2]的地址%v",&arr[0],&arr[1],&arr[2])
}

func TestArr02() {

	var arr [5]float64

	for i:=0;i<len(arr);i++{

		fmt.Println("请输入浮点数")
		fmt.Scanln(&arr[i])

	}

	for i := 0;i<len(arr);i++ {

		fmt.Printf("%v ",arr[i])
	}
}

func TestArr03() {

	arr := [3]int{1,2,3}

	//切片
	//var arr2 []int

	//index代表数组中索引的值，value代表索引对应的值
	for index,value := range arr{

		fmt.Printf("索引位置为%v,对应的值为%v\n",index,value)
	}
}

func TestSlice() {

	arr := [5]int{1,2,3,4,5}

	var slice []int
	//切片代表数组的一个引用
	slice = arr[1:3]

	fmt.Println(slice)
	//切片的容量动态变化
	fmt.Printf("%v\n",cap(slice))
	fmt.Printf("%v",len(slice))

}

/**
	内存分区：1.数组引用，切片和数组都可以操作数据
			2.make方式创建，只有切片可以操作数据
	注意：切片声明后不能直接使用，需要指定数组或者使用make分配内存
 */
func TestSlice2() {

	//切片的第一种初始化方式
	arr := [3]int{1,2,3}
	var slice []int
	slice = arr[1:2]

	for _,v:=range slice{

		fmt.Println(v)
	}

	//切片的第二种初始化方式,通过make的方式，指定长度和容量
	var slice2 = make([]int,2,4)
	for i := 0; i < len(slice2); i++ {

		slice2[i] = i
	}

	//切片的第三种方式
	var slice3 = []int{1,2,3}
	fmt.Printf("%v",slice3)
	fmt.Printf("%v",len(slice3))
	fmt.Printf("%v",cap(slice3))

	//切片还可以继续切片,引用数据类型，牵一发而动全身
	var sliceTemp []int
	sliceTemp = slice2[1:2]
	fmt.Println(sliceTemp)
}

func TestAutoAppend() {

	var slice = []int{1,2,3}

	//通过append方法向切片后面追加数值，实现自我增长
	slice = append(slice,11,12)

	//切片后面也可以追加切片
	slice = append(slice,slice...)

}

func TestSliceCopy() {

	var slice = []int{1,2,3,5}

	slice2 := make([]int,10)

	//将slice复制给slice2，两个切片是独立的，操作一个切片不会影响另一个切片
	copy(slice2,slice)

	//
	slice3 := make([]int,1)
	copy(slice3,slice)

}

func TestSliceStr() {

	//string底层为一个byte数组，所以string可以被切片使用
	str := "hilljay"

	//slice := str[:]

	//使用切片切 string类型的字符串后，不能直接更改，因为其实质还是一个string，如需更改，需要转成byte数组操作完后再转回string
	//slice[0] = "J"
	//fmt.Printf("%T\n",slice)//string

	//slice2 := []byte(str[:])
	////转为byte数组后修改则不会报错
	//slice2[0] = 'J'
	//str = string(slice2)
	//
	//fmt.Println(str)

	//如果有中文则将str转为rune数组
	slice3 := []rune(str)
	slice3[0] = '郭'
	str = string(slice3)
	fmt.Println(str)

}

func Fbn(n int) []uint64 {

	var fbn []uint64
	fbn = make([]uint64,n)

	for i := 0; i < n; i++ {

		if i==0||i==1 {

			fbn[i] = 1
		}else {

			fbn[i] = fbn[i-2]+fbn[i-1]
		}

	}

	return fbn
}

func TestDoubleArr() {

	var arr [2][3]int
	arr[0][0] = 1

}

func TestMap() {

	//map的使用方式一
	var map1 map[string]string
	map1 = make(map[string]string)

	//crud
	//增加
	map1["1"] = "123"
	//删除
	delete(map1,"1")
	//改
	map1["1"] = "456"
	//查
	val1 := map1["1"]
	fmt.Println(val1)

	//map的使用方式二
	map2 := make(map[string]string)

	//map的使用方式三
	var map3  = map[string]string{

		"1" : "hilljay",
		"2" : "jason",
	}

	fmt.Println(map1,map2,map3)
}

func TestMapRange() {

	var map1 map[string]string
	map1 = make(map[string]string)

	map1["1"] = "1"
	map1["2"] = "1"
	map1["3"] = "1"

	for key,val := range map1{

		fmt.Printf("key=%v,val=%v",key,val)
	}

}

/**
		map切片：切片在声明时需指定长度,而动态增长体现在append方法

 */
func TestSliceMap() {

	var mapSlice []map[string]string
	mapSlice  = make([]map[string]string,2)

	if mapSlice[0] == nil {

		mapSlice[0] = make(map[string]string)
		mapSlice[0]["1"] = "11"
		mapSlice[0]["2"] = "122"
	}

	if mapSlice[1] == nil {

		mapSlice[1] = make(map[string]string)
		mapSlice[0]["1"] = "11"
		mapSlice[0]["2"] = "122"
	}

	newMap := make(map[string]string)
	newMap["1"] = "22"

	mapSlice = append(mapSlice,newMap)
}

/**
	map排序
		//1.对key进行切片
		//2.对切片keys进行排序
		//3.对切片遍历取map中的值
 */
func SortMap() {

	intMap := make(map[int]int)
	intMap[9] = 2
	intMap[10] = 10
	intMap[1] = 1

	//声明一个切片
	var slice []int
	for k,_ := range intMap{

		slice = append(slice,k)
	}

	//对切片进行排序
	sort.Ints(slice)

	for index,_ := range slice{

		fmt.Println(intMap[slice[index]])
	}
}
