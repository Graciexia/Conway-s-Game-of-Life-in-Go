package main

import "fmt"


type Board struct {
	row,column int
	status [][]int
}


func Initial(row, column int) *Board {
	s := make([][]int, column)
	for i := range s {
		s[i] = make([]int, row)
	}
	return &Board{status: s, row: row, column: column}
}

func main() {
	l := Initial(4,5)
	fmt.Print("\x0c", l)
}