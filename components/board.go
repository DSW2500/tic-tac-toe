package components

//Board structure will house all the cells
type Board struct {
	Cells []*Cell
	Size  uint8
}

//CreateBoard :
func CreateBoard(size uint8) *Board {
	size *= size
	myArray := make([]*Cell, size)
	//Initially all cells will be blank:
	var i uint8
	for ; i < size; i++ {
		myArray[i] = NewCell()
	}
	return &Board{
		Cells: myArray,
		Size:  size,
	}
}
