package singleTesting

func AddUpper(n int) int{

	res := 0
	for i:=0;i<n-1;i++ {

		res += i
	}
	return res
}

func Diversion(j int, k int) int {

	return j/k
}