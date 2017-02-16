package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)


type Cell struct {
	status int
}
func (c *Cell) MakeAlive() {
	c.status = 1
}

func (c *Cell) Kill() {
	c.status = 0
}

func (c *Cell) IsAlive() bool {
	return (c.status == 1)
}

func (c *Cell) CopyCellData(from_cell Cell) {
	c.status = from_cell.status
}

//print cell format
func (c *Cell) ToString() string {
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

func (b *Board) Initialize(rows, columns int) {
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
			b.cells_current[row_index][col_index].Kill()
			b.cells_copy[row_index][col_index].Kill()
		}
	}
}

func (b *Board) SetLivingCells(cell_coords [][]int) {
	for _, coord := range cell_coords {
		b.makeCellAlive(coord[0], coord[1])
	}
}

func (b *Board) makeCellAlive(row_index, col_index int) {
	b.cells_current[row_index][col_index].MakeAlive()
}

func (b *Board) killCell(row_index, col_index int) {
	b.cells_current[row_index][col_index].Kill()
}

// always checks the copy
func (b *Board) liveSum(row_index, col_index int) int {
	count := 0
	start_row := row_index - 1; if start_row < 0 { start_row = 0 }
	end_row := row_index + 2; if end_row > b.rows { end_row = b.rows }
	start_col := col_index - 1; if start_col < 0 { start_col = 0 }
	end_col := col_index + 2; if end_col > b.cols { end_col = b.cols }
	for _, row := range b.cells_copy[start_row:end_row] {
		for _, cell := range row[start_col:end_col] {
			if cell.IsAlive() { count += 1 }
		}
	}
	return count
}

// checks the copy but sets the current cells
func (b *Board) makeLiveOrKill(row_index, col_index int) {
	total_live := b.liveSum(row_index, col_index)
	if total_live == 3 {
		b.makeCellAlive(row_index, col_index)
	} else if total_live != 4 {
		b.killCell(row_index, col_index)
	}
}

func (b *Board) dupCurrentToCopy() {
	for row_index := range b.cells_current {
		for col_index := range b.cells_current[row_index] {
			b.cells_copy[row_index][col_index].CopyCellData(b.cells_current[row_index][col_index])
		}
	}
}

func (b *Board) ChangeLife() {
	b.dupCurrentToCopy()
	for row_index := range b.cells_current {
		for col_index := range b.cells_current[row_index] {
			b.makeLiveOrKill(row_index, col_index)
		}
	}
}

func (b *Board) PrintBoard() {
	fmt.Println("+-" + strings.Repeat("--", b.cols) + "+")
	for _, row := range b.cells_current {
		fmt.Print("| ")
		for _, cell := range row {
			fmt.Print(cell.ToString() + " ")
		}
		fmt.Println("|")
	}
	fmt.Println("+-" + strings.Repeat("--", b.cols) + "+")
}

func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var board Board
	board.Initialize(10,16)
	//set up initial board configuration
	//Pentadecathlon
	board.SetLivingCells([][]int{{3,5}, {3,10}, {4,3}, {4,4}, {4,6}, {4,7}, {4,8}, {4,9}, {4,11}, {4,12}, {5,5}, {5,10}})
	//Glider
	//board.set_living_cells([][]int{{2,1}, {2,2}, {2,3}, {1,3}, {0,2}})
	clearScreen()
	board.PrintBoard()
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		board.ChangeLife()
		clearScreen()
		board.PrintBoard()
	}
}






