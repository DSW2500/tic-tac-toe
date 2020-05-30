package components

import (
	"fmt"
	"testing"
)

//testing if we get a correct NoMark Cell
func TestInitCell(t *testing.T) {
	expected := "-"
	actual := NewCell().Mark
	if expected != actual {
		t.Error("Not working")
	}
}

//Testing SetMark
func TestSetMark(t *testing.T) {
	expected2 := "X"
	actual2 := NewCell()
	_ = actual2.SetMark("X")
	if expected2 != actual2.Mark {
		t.Error("Not working")
	}

	fmt.Println(expected2, actual2.Mark)
}
func TestSlice(t *testing.T) {
	expected := []string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}
	var actual []string
	for i := 0; i < 9; i++ {
		newcell := NewCell()
		actual = append(actual, newcell.Mark)
	}

	fmt.Println(actual)
	fmt.Println(expected)
}
func TestGetMark(t *testing.T) {
	expected := "X"
	cell := &Cell{
		Mark: XMark,
	}
	actual := cell.GetMark()
	if expected != actual {
		t.Error("NOt working")
	}
}
