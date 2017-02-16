package main

import (
	"testing"
	"fmt"
)

func TestCell(t *testing.T) {
	c := Cell{1}
	if c.IsAlive() == false {
		t.Error("Expected true, got ", c.IsAlive())
	}

	c.Kill()
	if c.status != 0 {
		t.Error("Expected 0, got ", c.status)
	}

	c = Cell{1}
	c.MakeAlive()
	if c.status != 1 {
		t.Error("Expected 1, got ", c.status)
	}

	c_copy := Cell{0}
	c_copy.CopyCellData(c)
	if c_copy.status != 1 {
		t.Error("Expected 1, got ", c_copy.status)
	}
//test c.to_s()
	if  c.ToString()!= "*" {
		t.Error("Expected *, got ", c.ToString())
	}
	c = Cell{0}
	if  c.ToString()!= " " {
		t.Error("Expected *, got ", c.ToString())
	}
}


func TestBoard(t *testing.T) {
	var b Board
	s := "[[{0} {0} {0}] [{0} {0} {0}] [{0} {0} {0}]]"
	b.Initialize(3,3)
	if b.rows != 3 {
		t.Error("Expected 3, got ", b.rows)
	}
	if b.cols != 3 {
		t.Error("Expected 3, got ", b.cols)
	}

	if fmt.Sprint(b.cells_current) != s {
		t.Error("Expected", s, "got ", fmt.Sprint(b.cells_current))
	}

	if fmt.Sprint(b.cells_copy) != s {
		t.Error("Expected", s, "got ", fmt.Sprint(b.cells_current))
	}

	b.SetLivingCells([][]int{{0,0},{1,1},{2,2}})
	if b.cells_current[0][0].status != 1 {
		t.Error("Expected 1, got ", b.cells_current[0][0].status)
	}

	b.makeCellAlive(2,1)
	if  b.cells_current[2][1].status!= 1 {
		t.Error("Expected 1, got ", b.cells_current[2][1].status)
	}

	b.killCell(2,1)
	if  b.cells_current[2][1].status!= 0 {
		t.Error("Expected 1, got ", b.cells_current[2][2].status)
	}

	b.dupCurrentToCopy()
	if &b.cells_copy != &b.cells_current{
		t.Error("Expected true, got ", false)

	}
	//fmt.Println(&b.cells_copy)
	//fmt.Println(&b.cells_current)
	b.ChangeLife()
	if &b.cells_copy == &b.cells_current{
		t.Error("Expected false, got ", true)
	}

}

func TestBoardMakeLiveOrKill(t *testing.T) {
	var b Board
	b.Initialize(5,5)

	b.SetLivingCells([][]int{{0,3},{1,1},{1,2},{2,3}})
	b.dupCurrentToCopy()

//total live doesn't equal to 4 or 3 kill it
	if  b.liveSum(1, 1) != 2 {
		t.Error("Expected 2, got ", b.liveSum(1, 1))
	}

	b.makeLiveOrKill(1,1)
	if  b.cells_current[1][1].status!= 0 {
		t.Error("Expected 0, got ", b.cells_current[1][1].status)
	}

	if  b.liveSum(1, 0) != 1 {
		t.Error("Expected 1, got ", b.liveSum(1, 0))
	}

	b.makeLiveOrKill(1,0)
	if  b.cells_current[1][0].status!= 0 {
		t.Error("Expected 0, got ", b.cells_current[1][0].status)
	}

//total live equals to 4 keep the original live status
	if  b.liveSum(1, 2) != 4 {
		t.Error("Expected 4, got ", b.liveSum(1, 2))
	}

	b.makeLiveOrKill(1,2)
	if  b.cells_current[1][2].status!= 1 {
		t.Error("Expected 1, got ", b.cells_current[1][2].status)
	}
//total live equals to 3 set to alive
	if b.liveSum(1, 3) != 3 {
		t.Error("Expected 3, got ", b.liveSum(1, 3))
	}

	b.makeLiveOrKill(1,3)
	if  b.cells_current[1][3].status!= 1 {
		t.Error("Expected 1, got ", b.cells_current[1][3].status)
	}

}
