package gorountine

import "sync"

var UpperMap = make(map[int]int)
var lock sync.Mutex

func TestUpper(n int) {

	res :=0
	for i:=0;i<=n;i++{

		res *= i
	}

	//加锁以及释放锁
	lock.Lock()
	UpperMap[n] = res
	lock.Unlock()
}

//func GetMap() map[int]int {
//
//	return upperMap
//
//}
