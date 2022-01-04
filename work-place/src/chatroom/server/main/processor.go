package main

import (
	"fmt"
	"net"
	"work-place/src/chatroom/common"
	process2 "work-place/src/chatroom/server/process"
	"work-place/src/chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(message *common.Message) (err error) {

	switch message.Type {
	//选择登陆类型
	case common.LoginMessageType:
		userProcess := process2.UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.ServerProcessLogin(message)
	case common.RegisterType:

		userProcess := process2.UserProcess{Conn: this.Conn}
		err = userProcess.ServerProcessRegister(message)
		
	default:
		fmt.Println("处理类型不存在")
	}

	return
}

func (this *Processor) DoProcess() error {

	for {

		transFer := &utils.TransFer{Conn: this.Conn}
		requestMeg, err := transFer.HandleRequest()
		if err != nil {

			return err
		}

		err = this.serverProcessMes(&requestMeg)
		if err != nil {

			return err
		}
	}

}
