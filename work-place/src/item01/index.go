package item01

import "fmt"

func ShowInfo() {

	key := ""
	loop := true

	balance := 10000
	money := 0
	note := ""
	detail := "收支\t余额\t收入金额\t说明"

	outcome := 0
	char := ""

	//定义一个标识符
	flag := false
	for {
		fmt.Println("家庭记账系统")
		fmt.Println()
		fmt.Println(" 1.收支明细")
		fmt.Println(" 2.登记收入")
		fmt.Println(" 3.登记支出")
		fmt.Println(" 4.退出软件")
		fmt.Println("请选择（1-4）：")

		fmt.Scanln(&key)

		switch key {

		case "1":
			fmt.Println("当前收支明细记录\n")

			if flag {
				fmt.Println(detail)
			}else {
				fmt.Println("没有任何收支，来一笔")
			}

		case "2":
			fmt.Println("登记收入")
			fmt.Println("输出收入金额")
			fmt.Scanln(&money)
			fmt.Println("输入收入明细")
			fmt.Scanln(&note)
			balance += money
			detail += fmt.Sprintf("\n收入\t%v,\t%v,\t%v",balance,money,note)

			flag = true
		case "3":
			fmt.Println("登记支出")
			fmt.Println("输入支出金额")
			fmt.Scanln(&outcome)
			fmt.Println("输入支出明细")
			fmt.Scanln(&note)
			balance -= outcome
			detail += fmt.Sprintf("\n支出\t%v,\t%v,\t%v",balance,outcome,note)

			flag = true
		case "4":
			fmt.Println("退出")
			fmt.Println("确认要退出吗？y/n")

			fmt.Scanln(&char)

			if char=="y"{
				loop = false
			}else if char == "n"{
				break
			}

		default:
			fmt.Println("请输入正确的选项")
		}

		if (!loop){
			break
		}
	}

}
