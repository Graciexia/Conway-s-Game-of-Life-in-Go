package main

import (
	"testing"
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


