package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

//func (b *Board) put(x, y int, u string) // interface
//func (b *Board) get(x, y int) string    // interface

func input(n_m string) [2]int {
	var x [2]int
	s := strings.Split(n_m, ",")
	x[0], _ = strconv.Atoi(s[0])
	x[1], _ = strconv.Atoi(s[1])
	return x
}

func (b *Board) init() {
	for i := 0; i < 9; i++ {
		b.tokens[i] = 0
	}
}

func (b *Board) put(x, y int, u string) {

	//	if x >= 0 && x <= 2 && y >= 0 && y <= 2 {
	//エラー
	//	}
	//	if b.tokens[x+3*y] != 0 {
	//エラー
	//	}

	if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = -1
	}

	//return b*
}

func printBoard(b *Board) {
	for i := 0; i < 9; i++ {
		if b.tokens[i] == 0 {
			fmt.Print(".")
		} else if b.tokens[i] == 1 {
			fmt.Print("o")
		} else if b.tokens[i] == -1 {
			fmt.Print("x")
		}
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func checkWin(b *Board) int {
	if b.tokens[0] == b.tokens[1] && b.tokens[1] == b.tokens[2] { //横列
		if b.tokens[0] == 0 {
			return 0
		} else if b.tokens[0] == 1 {
			return 1
		} else if b.tokens[0] == -1 {
			return -1
		}
	}
	if b.tokens[3] == b.tokens[4] && b.tokens[4] == b.tokens[5] { //横列
		if b.tokens[3] == 0 {
			return 0
		} else if b.tokens[3] == 1 {
			return 1
		} else if b.tokens[3] == -1 {
			return -1
		}
	}
	if b.tokens[6] == b.tokens[7] && b.tokens[7] == b.tokens[8] { //横列
		if b.tokens[6] == 0 {
			return 0
		} else if b.tokens[6] == 1 {
			return 1
		} else if b.tokens[6] == -1 {
			return -1
		}
	}
	if b.tokens[0] == b.tokens[3] && b.tokens[3] == b.tokens[6] { //たて列
		if b.tokens[0] == 0 {
			return 0
		} else if b.tokens[0] == 1 {
			return 1
		} else if b.tokens[0] == -1 {
			return -1
		}
	}
	if b.tokens[1] == b.tokens[4] && b.tokens[4] == b.tokens[7] { //たて列
		if b.tokens[1] == 0 {
			return 0
		} else if b.tokens[1] == 1 {
			return 1
		} else if b.tokens[1] == -1 {
			return -1
		}
	}
	if b.tokens[2] == b.tokens[5] && b.tokens[5] == b.tokens[8] { //たて列
		if b.tokens[2] == 0 {
			return 0
		} else if b.tokens[2] == 1 {
			return 1
		} else if b.tokens[2] == -1 {
			return -1
		}
	}
	if b.tokens[0] == b.tokens[4] && b.tokens[4] == b.tokens[8] { //斜め
		if b.tokens[0] == 0 {
			return 0
		} else if b.tokens[0] == 1 {
			return 1
		} else if b.tokens[0] == -1 {
			return -1
		}
	}
	if b.tokens[2] == b.tokens[4] && b.tokens[4] == b.tokens[6] { //斜め
		if b.tokens[2] == 0 {
			return 0
		} else if b.tokens[2] == 1 {
			return 1
		} else if b.tokens[2] == -1 {
			return -1
		}
	}
	return 0
}

func main() {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var n_m string
	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			print("Player1 : input(x,y) ")
			fmt.Scan(&n_m)
			grid := input(n_m)
			b.put(grid[0], grid[1], "o")
			printBoard(b)
		}
		if i%2 == 1 {
			print("Player2 : input(x,y) ")
			fmt.Scan(&n_m)
			grid := input(n_m)
			b.put(grid[0], grid[1], "x")
			printBoard(b)
		}
		if checkWin(b) == 1 {
			fmt.Println("Player1 won")
			break
		} else if checkWin(b) == -1 {
			fmt.Println("Player2 won")
			break
		}
		if i == 8 {
			fmt.Println("Draw")
		}
	}

}

/*
func main() {
	//		var board Board
	//		board = ['.','.','.']
	var grid [2]int
	var n_m string
	print("Player1 :input (x,y) ")
	fmt.Scan(&n_m)

	grid = input(n_m)

	var b Board
	//b := &Board{tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	b.init()
	b.put(grid[0], grid[1], "o")
	//printBoard(b)
}
*/
