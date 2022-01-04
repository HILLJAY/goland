package process

import (
	"encoding/json"
	"fmt"
	"net"
	"work-place/src/chatroom/common"
	"work-place/src/chatroom/server/model"
	"work-place/src/chatroom/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessRegister(mesg *common.Message) (err error) {

	return
}

func (this *UserProcess) ServerProcessLogin(mesg *common.Message) error {

	//1.首先获取传入信息中的Data部分，将其反序列化为一个 common.LoginMessage 类型的Message
	var loginMes common.LoginMessage
	err := json.Unmarshal([]byte(mesg.Data), &loginMes)
	if err != nil {

		return err
	}

	//2.再声明一个resMes,最终传递的是这个Message
	var resMeg common.Message
	resMeg.Type = common.LoginResType

	//3.声明一个 common.ResponseMessage
	var loginResMessage common.ResponseMessage

	usr, err := model.SingleUserDao.ChargeLogin(loginMes.UserId, loginMes.UserPwd)
	fmt.Println(usr.UserId, "登陆成功")
	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {

			loginResMessage.Code = 500
			loginResMessage.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {

			loginResMessage.Code = 403
			loginResMessage.Error = err.Error()
		} else {

			loginResMessage.Code = 505
			loginResMessage.Error = "服务器内部错误"
		}
	} else {

		loginResMessage.Code = 200
		fmt.Println("登陆成功")
	}

	//4.将loginResMessage序列化后赋值给resMes的Data部分
	bytes, err := json.Marshal(loginResMessage)
	if err != nil {
		return err
	}
	resMeg.Data = string(bytes)

	//5.将resMes序列化，然后进行传输
	marshal, err := json.Marshal(resMeg)
	if err != nil {
		return err
	}

	//6.将返回的数据写入连接中，传输给客户端
	fer := &utils.TransFer{
		Conn: this.Conn,
	}
	err = fer.HandleResponse(marshal)
	if err != nil {
		return err
	}

	return err
}
