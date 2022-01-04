package main

import (
	"fmt"
	"work-place/src/item02"
)

func main() {

	mainView := mainView{}
	mainView.customerService = item02.UserServiceFactory()
	mainView.showIndex()
}

type mainView struct {

	customerService *item02.UserService

}

func (cv *mainView) updateUser(id int, name string, address string, email string) bool{

	return cv.updateUser(id,name,address,email)
}

func (cv *mainView) list() {

	customers := cv.customerService.List()

	fmt.Println("--------客户列表--------")
	fmt.Println("编号\t姓名\t地址\t邮箱")
	for i := 0; i < len(customers); i++ {

		customers[i].GetInfo()
	}

}

func (cv *mainView) addUser(id int, name string, address string, email string) {

	cv.customerService.AddUser(id,name,address,email)
}

func (cv *mainView) delUser(id int) bool{

	return cv.customerService.DelUser(id)
}

func (cv *mainView) showIndex() {

	key := ""
	loop := true

	for {
		fmt.Println("------------客户信息管理系统------------")
		fmt.Println("			1.添加用户")
		fmt.Println("			2.删除用户")
		fmt.Println("			3.修改用户")
		fmt.Println("			4.用户列表")
		fmt.Println("			5.退出")
		fmt.Println("			请输入(1-5)")

		fmt.Scanln(&key)

		switch key {

		case "1":
			fmt.Println("添加用户")
			fmt.Println("输入学号")
			id := 1
			fmt.Scanln(&id)
			fmt.Println("输入姓名")
			name := ""
			fmt.Scanln(&name)
			fmt.Println("输入地址")
			address := ""
			fmt.Scanln(&address)
			fmt.Println("输入邮箱")
			email := ""
			fmt.Scanln(&email)
			cv.addUser(id,name,address,email)
		case "2":
			fmt.Println("删除用户")
			fmt.Println("输入删除用户的id")
			id := 0
			fmt.Scanln(&id)
			if !cv.delUser(id){
				fmt.Println("无此用户")
			}else {
				cv.delUser(id)
				fmt.Println("删除成功")
			}
		case "3":
			fmt.Println("修改用户")
			fmt.Println("输入要修改用户的id")
			id := 0
			fmt.Scanln(&id)
			fmt.Println("用户信息如下")
			fmt.Printf("id:%V",cv.customerService.GetCustomers()[id].GetId())

		case "4":
			fmt.Println("用户列表")
			cv.list()
		case "5":
			fmt.Println("退出")
			loop = false
		}

		if !loop {
			break
		}

	}

}

