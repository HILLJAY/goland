package testSign

import "fmt"

func Main() {

	testNumOfPass()
}

func testSum() {

	var sum int = 0

	//可以指定标签以便于break 跳出指定标签的循环
	test01:
	for i := 0;i<=100;i++ {

		sum += i
		if sum > 20 {

			fmt.Println(i)
			break test01
 		}
	}

}

func testCount() {

	var positiveCount int
	var negativeCount int
	var num int

	for {
		fmt.Println("输入数值：")
		fmt.Scanln(&num)

		if num == 0 {

			break
		}

		if num>0 {
			positiveCount++
			continue
		}

		negativeCount++
	}

	fmt.Printf("正数的个数%v,负数的个数%v",positiveCount,negativeCount)
}

func testNumOfPass() {

	var money float64 = 100000
	var count int = 0

	for {

		if money > 50000 {

			money = money - (money * 0.05)
			count++
		}else {

			money = money - 1000
			count ++
		}

		if money<0 {

			count --
			break
			}
		}

		fmt.Println(count)
}

func testGoto() {

	var money float64 = 100000
	var count int = 0

	for {

		if money > 50000 {

			money = money - (money * 0.05)
			count++
		}else {

			money = money - 1000
			count ++
		}

		if money<0 {	

			count --
			break
		}
	}

	fmt.Println(count)

}