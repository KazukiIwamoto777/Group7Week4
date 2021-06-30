package main

import "testing"

func TestEx01(t *testing.T) {
	expect := [2]int{1, 1}
	result := input("1,1")
	if result != expect {
		t.Error("Test01 is failed")
	}
}
