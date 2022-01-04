package process

import (
	"fmt"
	"os"
)

func ShowMenu() {

	fmt.Println("-------------欢迎登陆----------------")
	fmt.Println("-------------1.显示用户列表----------------")
	fmt.Println("-------------2.发送消息----------------")
	fmt.Println("-------------3.信息列表----------------")
	fmt.Println("-------------4.退出系统----------------")

	fmt.Println("\t\t\t 请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:

		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:

		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误")
	}
}
