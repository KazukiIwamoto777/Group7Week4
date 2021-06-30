package main

import (
	"fmt"
	//"strings"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

//func (b *Board) put(x, y int, u string) // interface
//func (b *Board) get(x, y int) string    // interface

func input(n_m string) int {
	//	s := strings.Split(n_m, ",")
	x := 1
	// y := 2
	return x
}

/*
func (b *Board) put(x, y int, u string) {

	if x >= 0 && x <= 2 && y >= 0 && y <= 2 {
		//エラー
	}
	if b.tokens[x+3*y] != 0 {
		//エラー
	}

	if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = -1
	}

	return b
}
*/
func main() {
	//		var board Board
	//		board = ['.','.','.']
	var n_m string
	print("Player1 :input (x,y) ")
	fmt.Scan(&n_m)
	fmt.Println(input(n_m))
}
