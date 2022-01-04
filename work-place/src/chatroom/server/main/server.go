package main

import (
	"fmt"
	"net"
	"time"
	"work-place/src/chatroom/server/model"
)

/**
server
*/
func main() {

	//服务器开始的时候就建立Redis连接
	initPool("192.168.1.10:6379", 8, 10, 300*time.Second)
	//初始化UserDao
	initUserDao()

	listen, err := net.Listen("tcp", "127.0.0.1:8889")

	if err != nil {

		fmt.Println("error", err)
	}

	defer listen.Close()

	for {
		con, err := listen.Accept()
		if err != nil {

			fmt.Println(err)
		} else {

			go process(con)
		}
	}

}

func initUserDao() {

	model.SingleUserDao = model.UserDaoFactory(pool)
}

func process(con net.Conn) {

	defer con.Close()

	processor := &Processor{Conn: con}

	err := processor.DoProcess()
	if err != nil {

		fmt.Println(err)
		return
	}
}

//func handleRequest(conn net.Conn) (mes common.Message, err error) {
//
//	buf := make([]byte, 1024)
//
//	_, err = conn.Read(buf[:4])
//	if err != nil {
//
//		return
//	}
//
//	//获取packageLength
//	var pkgLength uint32
//	pkgLength = binary.BigEndian.Uint32(buf[:4])
//
//	//根据packageLength获取信息
//	n, err := conn.Read(buf[:pkgLength])
//	if uint32(n) != pkgLength || err != nil {
//
//		return
//	}
//
//	//序列化信息，封装为一个mes,注意这里一定要穿地址值，如果没传则默认为值拷贝，无法修改message
//	err = json.Unmarshal(buf[:pkgLength], &mes)
//	if err != nil {
//		return
//	}
//
//	return
//}
//
//func handleResponse(conn net.Conn, mesBytes []byte) (err error) {
//
//	//1.先获取数据长度
//	var pkgLen uint32
//	pkgLen = uint32(len(mesBytes))
//
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[:4], pkgLen)
//
//	//2.先发送数据长度
//	n, err := conn.Write(buf[:4])
//	fmt.Println(n)
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write(bytes) fail", err)
//		return
//	}
//
//	//3.发送数据本身
//	_, err = conn.Write(mesBytes)
//
//	if err != nil {
//		return err
//	}
//
//	return
//}
//
///**
//编写一个serverProcessMes，根据客户端传入的消息类型，调用不同的函数处理
//*/
//func serverProcessMes(conn net.Conn, message *common.Message) (err error) {
//
//	switch message.Type {
//	//选择登陆类型
//	case common.LoginMessageType:
//		err = serverProcessLogin(conn, message)
//	case common.RegisterType:
//
//	default:
//		fmt.Println("处理类型不存在")
//	}
//
//	return
//}
//
///**
//处理登陆类型请求
//*/
//func serverProcessLogin(conn net.Conn, mesg *common.Message) error {
//
//	//1.首先获取传入信息中的Data部分，将其反序列化为一个 common.LoginMessage 类型的Message
//	var loginMes common.LoginMessage
//	err := json.Unmarshal([]byte(mesg.Data), &loginMes)
//	if err != nil {
//
//		return err
//	}
//
//	//2.再声明一个resMes,最终传递的是这个Message
//	var resMeg common.Message
//	resMeg.Type = common.LoginResType
//
//	//3.声明一个 common.ResponseMessage
//	var loginResMessage common.ResponseMessage
//
//	if loginMes.UserId == 1 && loginMes.UserPwd == "123456" {
//
//		loginResMessage.Code = 200
//	} else {
//
//		loginResMessage.Code = 500
//		loginResMessage.Error = "该用户不存在"
//	}
//
//	//4.将loginResMessage序列化后赋值给resMes的Data部分
//	bytes, err := json.Marshal(loginResMessage)
//	if err != nil {
//		return err
//	}
//	resMeg.Data = string(bytes)
//
//	//5.将resMes序列化，然后进行传输
//	marshal, err := json.Marshal(resMeg)
//	if err != nil {
//		return err
//	}
//
//	//6.将返回的数据写入连接中，传输给客户端
//	err = handleResponse(conn, marshal)
//	if err != nil {
//		return err
//	}
//
//	return err
//}
