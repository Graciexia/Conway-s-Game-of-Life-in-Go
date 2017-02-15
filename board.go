package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)


type Cell struct {
	status int
}
func (c *Cell) make_alive() {
	c.status = 1
}
func (c *Cell) kill() {
	c.status = 0
}
func (c *Cell) is_alive() bool {
	return (c.status == 1)
}
func (c *Cell) to_s() string {
	if c.status == 1 {
		return "*"
	}
	return " "
}


type Board struct {
	rows    int
	cols    int
	cells   [][]Cell
}
func (b *Board) initialize(rows, columns int) {
	b.rows = rows
	b.cols = columns
	b.cells = make([][]Cell, b.rows)
	for i := range b.cells {
		b.cells[i] = make([]Cell, b.cols)
	}
	for _, row := range b.cells {
		for col_ind := range row {
			//row[col_ind].kill()
			row[col_ind].make_alive()
		}
	}
}
func (b *Board) live_sum(row_index, col_index int) int {
	count := 0
	start_row := row_index - 1; if start_row < 0 { start_row = 0 }
	end_row := row_index + 2; if end_row > b.rows { end_row = b.rows }
	start_col := col_index - 1; if start_col < 0 { start_col = 0 }
	end_col := col_index + 2; if end_col > b.cols { end_col = b.cols }
	for _, row := range b.cells[start_row:end_row] {
		for _, cell := range row[start_col:end_col] {
			if cell.is_alive() { count += 1 }
		}
	}
	return count
}
func (b *Board) print_board() {
	clear_screen()
	fmt.Println("+-" + strings.Repeat("--", b.cols) + "+")
	for _, row := range b.cells {
		fmt.Print("| ")
		for _, cell := range row {
			fmt.Print(cell.to_s() + " ")
		}
		fmt.Println("|")
	}
	fmt.Println("+-" + strings.Repeat("--", b.cols) + "+")
}

func (b *Board) live_or_die(row_index, col_index int) {
	//total_lilve = func field_sum(board, row_index, col_index)

}

func clear_screen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var board Board
	board.initialize(10,10)
	clear_screen()
	board.print_board()
	fmt.Println(board.live_sum(3,3))
}

