package main

import "testing"

func TestEx01(t *testing.T) {
	expect := [2]int{1, 1}
	result := input("1,1")
	if result != expect {
		t.Error("Test01 is failed")
	}
}

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expect := []int{0, 0, 0, 0, 1, 0, 0, 0, 0}
	b.put(1, 1, "o")

	if expect[0] != b.tokens[0] {
		t.Error("Test02 is failed")
	}
	if expect[1] != b.tokens[1] {
		t.Error("Test02 is failed")
	}
	if expect[2] != b.tokens[2] {
		t.Error("Test02 is failed")
	}
	if expect[3] != b.tokens[3] {
		t.Error("Test02 is failed")
	}
	if expect[4] != b.tokens[4] {
		t.Error("Test02 is failed")
	}
	if expect[5] != b.tokens[5] {
		t.Error("Test02 is failed")
	}
	if expect[6] != b.tokens[6] {
		t.Error("Test02 is failed")
	}
	if expect[7] != b.tokens[7] {
		t.Error("Test02 is failed")
	}
	if expect[8] != b.tokens[8] {
		t.Error("Test02 is failed")
	}

}

/*
func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0,0,0,0,0,0,0,0,0}
	}
	b.put(1,1, "o")
	if b.get(1,1) != "o" {
		t.Errorf("....")
	}
}
*/
