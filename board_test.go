package main

import (
	"testing"
	"fmt"
)

func TestCell(t *testing.T) {
	c := Cell{1}
	if c.is_alive() == false {
		t.Error("Expected true, got ", c.is_alive())
	}

	c.kill()
	if c.status != 0 {
		t.Error("Expected 0, got ", c.status)
	}

	c = Cell{1}
	c.make_alive()
	if c.status != 1 {
		t.Error("Expected 1, got ", c.status)
	}

	c_copy := Cell{0}
	c_copy.copy_cell_data(c)
	if c_copy.status != 1 {
		t.Error("Expected 1, got ", c_copy.status)
	}
//test c.to_s()
	if  c.to_s()!= "*" {
		t.Error("Expected *, got ", c.to_s())
	}
	c = Cell{0}
	if  c.to_s()!= " " {
		t.Error("Expected *, got ", c.to_s())
	}
}


func TestBoard(t *testing.T) {
	var b Board
	s := "[[{0} {0} {0}] [{0} {0} {0}] [{0} {0} {0}]]"
	b.initialize(3,3)
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

	b.set_living_cells([][]int{{0,0},{1,1},{2,2}})
	if b.cells_current[0][0].status != 1 {
		t.Error("Expected 1, got ", b.cells_current[0][0].status)
	}

	b.make_cell_alive(2,1)
	if  b.cells_current[2][1].status!= 1 {
		t.Error("Expected 1, got ", b.cells_current[2][1].status)
	}

	b.kill_cell(2,1)
	if  b.cells_current[2][1].status!= 0 {
		t.Error("Expected 1, got ", b.cells_current[2][2].status)
	}

	//b.change_life()
	//if &b.cells_copy == &b.cells_current{
	//	t.Error("Expected false, got ", true)
	//
	//}
	//
	//b.copy_cells()
	//if *b.cells_copy != *b.cells_current{
	//	t.Error("Expected true, got ", false)
	//
	//}
}

func TestBoardMakeLiveOrKill(t *testing.T) {
	var b Board
	b.initialize(5,5)

	b.set_living_cells([][]int{{0,3},{1,1},{1,2},{2,3}})

//total live doesnt equal to 4 or 3 kill it
	if  b.live_sum(1, 1) != 2 {
		t.Error("Expected 2, got ", b.live_sum(0, 0))
	}

	b.make_live_or_kill(1,1)
	if  b.cells_current[1][1].status!= 0 {
		t.Error("Expected 0, got ", b.cells_current[1][1].status)
	}

	if  b.live_sum(1, 0) != 1 {
		t.Error("Expected 1, got ", b.live_sum(1, 0))
	}

	b.make_live_or_kill(1,0)
	if  b.cells_current[1][0].status!= 0 {
		t.Error("Expected 0, got ", b.cells_current[1][0].status)
	}

//total live equals to 4 keep the original live status
	if  b.live_sum(1, 2) != 4 {
		t.Error("Expected 4, got ", b.live_sum(1, 2))
	}

	b.make_live_or_kill(1,2)
	if  b.cells_current[1][2].status!= 1 {
		t.Error("Expected 1, got ", b.cells_current[1][2].status)
	}
//total live equals to 3 set to alive
	if b.live_sum(1, 3) != 3 {
		t.Error("Expected 3, got ", b.live_sum(1, 3))
	}

	b.make_live_or_kill(1,3)
	if  b.cells_current[1][3].status!= 1 {
		t.Error("Expected 1, got ", b.cells_current[1][3].status)
	}

}
