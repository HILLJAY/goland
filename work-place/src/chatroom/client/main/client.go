package main

import (
	"fmt"
	"os"
	process2 "work-place/src/chatroom/client/process"
)

func main() {

	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true
	var userId int
	var userPwd string
	var userName string

	for true {
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			//Login(userId, userPwd)
			up := process2.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("请输入用户信息")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的姓名")
			fmt.Scanf("%s\n", &userName)

			up := &process2.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

	}

}
