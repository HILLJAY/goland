package file

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Test01() {

	file, err := os.Open("hello.txt")

	if err==nil {

		fmt.Println(file)
	}else {

		fmt.Println(err)
	}

	erro := file.Close()
	if erro!=nil {

		fmt.Println(erro)
	}
}

func Test02() {

	//通过os.Open打开文件，返回一个file指针
	file, _ := os.Open("hello.txt")

	//使用defer关闭文件流
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	//通过缓冲区读取file
	reader := bufio.NewReader(file)

	//循环读取文本
	for {

		// delim : 以什么为分隔符
		str,err := reader.ReadString('\n')
		fmt.Printf(str)

		// io.EOF 代表当到达文本最后一行时break
		if err==io.EOF {
			break
		}

	}

	fmt.Println("\n读取结束")
}

func Test03() {

	//通过 ioutil.ReadFile直接读取指定文件，方法内部把open 和 close方法封装了，适用于文件较小的文件读取
	file, err := ioutil.ReadFile("hello.txt")

	if err!=nil{

		fmt.Println(err)
	}

	fmt.Println(string(file))
}

func Test04()  {

	var filePath string
	filePath = "d:/abc.txt"

	//文件路径，模式选择创建和写入，Linux/Unix下的文件权限设置
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)

	if err!=nil{

		fmt.Println(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	str := "hello garden\n"

	for i:=0;i<5;i++{

		writer.WriteString(str)
	}

	//在使用Writer缓冲流写入数据时，数据是先写入到缓冲中，需要刷新才会写入到文件中
	writer.Flush()
}

func Test05() {

	//先读后写入指定目录下

	filePath1 := "d:/abc.txt"
	filePath2 := "d:/kkk.txt"

	file,errs := ioutil.ReadFile(filePath1)

	if errs != nil {

		fmt.Printf("%v\n",errs)
	}

	err := ioutil.WriteFile(filePath2,file,0666)

	if err!=nil {

		fmt.Printf("%v\n",err)
	}
}

func chargeFile(file string) (bool,error) {

	_, err := os.Stat(file)
	if err==nil {

		return true,nil
	}else if os.IsExist(err){

		return false,nil
	}

	return false,err
}

func copyFile(des string, src string) (writtern int64, err error) {

	file, err := os.Open(src)
	defer file.Close()

	if err!=nil {

		fmt.Printf("%v",err)
		return
	}
	reader := bufio.NewReader(file)

	//打开另一个文件
	openFile, errs := os.OpenFile(des, os.O_CREATE|os.O_WRONLY, 0666)
	defer openFile.Close()

	if errs !=nil {

		fmt.Printf("%v",errs)
		return
	}

	writer := bufio.NewWriter(openFile)
	return io.Copy(writer,reader)
}

func Test06() {

	fmt.Printf("命令行的参数有%v",len(os.Args))
	for i,v := range os.Args{

		fmt.Printf("agrs[%v]=%v\n",i,v)
	}
}

func TestFlag() {

	var usr string
	var pwd string
	var host string
	var port int

	flag.StringVar(&usr,"u","","用户名默认为空")
	flag.StringVar(&pwd,"","","密码默认为空")
	flag.StringVar(&host,"h","localhost","主机名，默认为localhost")
	flag.IntVar(&port,"port",3306,"端口号，默认3306")

	//解析
	flag.Parse()

	fmt.Printf("name=%v",usr)
}