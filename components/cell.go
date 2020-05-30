package components

import (
	"errors"
)

//Cell :
type Cell struct {
	Mark string
}

//Enums
const (
	NoMark = "-" //Initial value is nil
	XMark  = "X"
	OMark  = "O"
)

//NewCell :
func NewCell() *Cell {
	cell := &Cell{Mark: NoMark}
	return cell
}

//SetMark : user sets a mark on the cell
func (cell *Cell) SetMark(mark string) error {
	if cell.Mark != "-" {
		err := errors.New("writing to a non-empty cell")
		// setMark(mark,cell)
		return err
	}

	cell.Mark = mark
	return nil
}

//GetMark :
func (cell *Cell) GetMark() string {
	return cell.Mark

}
