package components

import "testing"

//Testing if a newly created board has all cells marked NoMark  "-"
func TestBoard(t *testing.T) {
	var size uint8 = 9
	expected := &Board{
		Cells: []*Cell{
			{Mark: NoMark},
			{Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}},
		Size: size,
	}
	actual := CreateBoard(9)
	var i uint8
	for i = 0; i < size; i++ {
		if actual.Cells[i].Mark != expected.Cells[i].Mark {
			t.Error("Not working")
		}
	}
}

//Testing if actual board has correct number of cells
func TestTotalCount(t *testing.T) {
	var count uint8 = 0
	var size uint8 = 9
	actual := CreateBoard(9)
	var i uint8
	for i = 0; i < size; i++ {
		if actual.Cells[i].Mark == "" {
			t.Error("Not working")
		}
		count++
	}
	if count != size {
		t.Error("Not working")
	}

}
