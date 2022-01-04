package singleTesting

import "testing"

func TestAddUpper(t *testing.T) {

	res := AddUpper(10)
	if res!=36{

		t.Fatalf("期望值%v,实际值%v",36,res)
	}

	t.Logf("This is a log for the calculate")
}

func TestDiversion(t *testing.T) {

	res := Diversion(8,2)

	t.Logf("res :%v",res)

}