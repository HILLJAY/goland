package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"work-place/src/chatroom/client/util"
	"work-place/src/chatroom/common"
)

type UserProcess struct {
}

func (this *UserProcess) Register(id int, pwd string, name string) (err error) {

	conn, err := net.Dial("tcp", "127.0.0.1:8889")

	if err != nil {

		fmt.Println(err)
	}
	//延时关闭
	defer conn.Close()
	//初始化message
	var message common.Message
	message.Type = common.RegisterType

	var registerMes common.RegisterMessage
	registerMes.User.UserId = id
	registerMes.User.UserPwd = pwd
	registerMes.User.UserName = name

	bytes, _ := json.Marshal(registerMes)
	message.Data = string(bytes)

	//序列化message
	mesgJson, err := json.Marshal(message)
	if err != nil {

		fmt.Println(err)
	}

	var pkgLen uint32
	pkgLen = uint32(len(mesgJson))

	var length [4]byte
	binary.BigEndian.PutUint32(length[0:4], pkgLen)

	n, err := conn.Write(length[:4])

	if n != 4 || err != nil {

		fmt.Println("传输异常")
		return
	}

	tf := util.TransFer{Conn: conn}
	err = tf.HandleResponse(mesgJson)
	if err != nil {
		return
	}

	//处理服务端给客户端的返回信息
	transFer := &util.TransFer{Conn: conn}
	resMes, err := transFer.HandleRequest()

	if err != nil {
		fmt.Println(err)
		return
	}

	//将其Data值反序列化为一个responseMes
	var responseMes common.ResponseMessage
	err = json.Unmarshal([]byte(resMes.Data), &responseMes)
	if err != nil {
		return
	}

	if responseMes.Code == 200 {
		fmt.Println("注册成功")
		os.Exit(0)
	} else if responseMes.Code == 500 {

		fmt.Println(responseMes.Error)
		return
	}

	return
}

func (this *UserProcess) Login(userId int, userPwd string) {

	conn, err := net.Dial("tcp", "127.0.0.1:8889")

	if err != nil {

		fmt.Println(err)
	}
	//延时关闭
	defer conn.Close()
	//初始化message
	var message common.Message
	message.Type = common.LoginMessageType

	var data common.LoginMessage
	data.UserId = userId
	data.UserPwd = userPwd

	bytes, _ := json.Marshal(data)
	message.Data = string(bytes)

	//序列化message
	mesgJson, err := json.Marshal(message)
	if err != nil {

		fmt.Println(err)
	}

	//传输的过程中需要的是一个byte[]切片，所以需要将数据长度转换为一个uint类型的变量
	var pkgLen uint32
	pkgLen = uint32(len(mesgJson))

	var length [4]byte
	binary.BigEndian.PutUint32(length[0:4], pkgLen)

	n, err := conn.Write(length[:4])

	if n != 4 || err != nil {

		fmt.Println("传输异常")
		return
	}

	_, err = conn.Write(mesgJson)
	if err != nil {

		fmt.Println(err)
		return
	}

	//处理服务端给客户端的返回信息
	transFer := &util.TransFer{Conn: conn}
	resMes, err := transFer.HandleRequest()

	if err != nil {
		fmt.Println(err)
		return
	}

	//将其Data值反序列化为一个responseMes
	var responseMes common.ResponseMessage
	err = json.Unmarshal([]byte(resMes.Data), &responseMes)
	if err != nil {
		return
	}

	if responseMes.Code == 200 {
		fmt.Println("登陆成功")

		//登陆成功，显示菜单
		for {
			ShowMenu()
		}
	} else if responseMes.Code == 500 {

		fmt.Println(responseMes.Error)
		return
	}
}
