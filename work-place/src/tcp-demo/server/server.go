package server

import (
	"fmt"
	"net"
)

func process(con net.Conn) {

	defer con.Close()

	for {

		buf := make([]byte,1024)
		//fmt.Printf("服务器等待客户端%s发送消息",con.RemoteAddr().String())

		n, err := con.Read(buf)
		if err!=nil {

			fmt.Println("处理错误")
			return
		}else {

			fmt.Print(string(buf[:n]))
		}

	}
}

func Server() {

	listen, err := net.Listen("tcp", "127.0.0.1:8888")

	if err!=nil{

		fmt.Println("错误为",err)
	}

	defer listen.Close()

	for {
		//循环等待连接
		con, err := listen.Accept()
		if err!=nil{
			fmt.Println("连接错误")
		}else {

			fmt.Println("客户端连接成功",con.RemoteAddr().String())
			go process(con)
		}
	}
}
