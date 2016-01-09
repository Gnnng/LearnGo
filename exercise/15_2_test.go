package stack

import "testing"

func TestPushAndPop(t *testing.T) {
	var a Stack
	a.Push(1)
	if a[0] != 1 {
		t.Log("Not successfully pushed")
		t.Fail()
	} else {
		t.Log("Pushed")
	}
	x := a.Pop()
	if x != 1 || len(a) != 0 {
		t.Log("Not successfully poped")
		t.Fail()
	} else {
		t.Log("Poped")
	}
}
