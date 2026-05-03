package goaddsample

import "testing"

func Test_Add(t *testing.T){
	result:= Add[int](3,4)
	if result != 7 {
		t.Error("incorrect result: expected 7, got",result)
	}
	resultFloat:= Add[float64](3.5,4.2)
	if resultFloat != 7.7 {
		t.Error("incorrect result: expected 7.7, got", resultFloat)
	}
}
