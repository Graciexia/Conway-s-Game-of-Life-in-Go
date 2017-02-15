package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	//"time"
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

func (c *Cell) copy_cell_data(from_cell Cell) {
	c.status = from_cell.status
}

//print cell format
func (c *Cell) to_s() string {
	if c.status == 1 {
		return "*"
	}
	return " "
}


type Board struct {
	rows          int
	cols          int
	cells_current [][]Cell
	cells_copy    [][]Cell
}

func (b *Board) initialize(rows, columns int) {
	b.rows = rows
	b.cols = columns
	b.cells_current = make([][]Cell, b.rows)
	b.cells_copy = make([][]Cell, b.rows)
	for i := range b.cells_current {
		b.cells_current[i] = make([]Cell, b.cols)
		b.cells_copy[i] = make([]Cell, b.cols)
	}
	for row_index := range b.cells_current {
		for col_index := range b.cells_current[row_index] {
			b.cells_current[row_index][col_index].kill()
			b.cells_copy[row_index][col_index].kill()
		}
	}
}

func (b *Board) set_living_cells(cell_coords [][]int) {
	for _, coord := range cell_coords {
		b.make_cell_alive(coord[0], coord[1])
	}
}

func (b *Board) make_cell_alive(row_index, col_index int) {
	b.cells_current[row_index][col_index].make_alive()
}

func (b *Board) kill_cell(row_index, col_index int) {
	b.cells_current[row_index][col_index].kill()
}

// always checks the copy
func (b *Board) live_sum(row_index, col_index int) int {
	count := 0
	start_row := row_index - 1; if start_row < 0 { start_row = 0 }
	end_row := row_index + 2; if end_row > b.rows { end_row = b.rows }
	start_col := col_index - 1; if start_col < 0 { start_col = 0 }
	end_col := col_index + 2; if end_col > b.cols { end_col = b.cols }
	for _, row := range b.cells_copy[start_row:end_row] {
		for _, cell := range row[start_col:end_col] {
			if cell.is_alive() { count += 1 }
		}
	}
	return count
}

// checks the copy but sets the current cells
func (b *Board) make_live_or_kill(row_index, col_index int) {
	total_live := b.live_sum(row_index, col_index)
	if total_live == 3 {
		b.make_cell_alive(row_index, col_index)
	} else if total_live != 4 {
		b.kill_cell(row_index, col_index)
	}
}


func (b *Board) change_life() {
	b.copy_cells()
	for row_index := range b.cells_copy {
		for col_index := range b.cells_copy[row_index] {
			b.make_live_or_kill(row_index, col_index)
		}
	}
}


func (b *Board) copy_cells() {
	for row_index := range b.cells_current {
		for col_index := range b.cells_current[row_index] {
			b.cells_copy[row_index][col_index].copy_cell_data(b.cells_current[row_index][col_index])
		}
	}
}

func (b *Board) print_board() {
	clear_screen()
	fmt.Println("+-" + strings.Repeat("--", b.cols) + "+")
	for _, row := range b.cells_current {
		fmt.Print("| ")
		for _, cell := range row {
			fmt.Print(cell.to_s() + " ")
		}
		fmt.Println("|")
	}
	fmt.Println("+-" + strings.Repeat("--", b.cols) + "+")
}

func clear_screen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var board Board
	board.initialize(10,16)
	 //set up initial board configuration
	 //Pentadecathlon
	//board.set_living_cells([][]int{{3,5}, {3,10}, {4,3}, {4,4}, {4,6}, {4,7}, {4,8}, {4,9}, {4,11}, {4,12}, {5,5}, {5,10}})
	 //Glider
	board.set_living_cells([][]int{{2,1}, {2,2}, {2,3}, {1,3}, {0,2}})
	//board.print_board()
	//for i := 0; i < 100; i++ {
	//	time.Sleep(10 * time.Millisecond)
	//	board.change_life()
	//	board.print_board()
	//}

	//var s = fmt.Sprint(board.cells_current)
	//fmt.Println(s)
	//board.set_living_cells([][]int{{0,0},{1,1},{2,2}})
	////board.make_cell_alive(1,1)
	//a := board.live_sum(0,0)
	//fmt.Println(board.cells_current)
	fmt.Println(board.live_sum(2,3))

}






