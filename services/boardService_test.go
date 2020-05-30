package services

import (
	"fmt"
	"testing"
	"ticgame/components"
)

func TestNewBoardService(t *testing.T) {
	expected := &BoardService{
		Board: components.CreateBoard(3),
	}
	NoMark := components.NoMark
	var size uint8
	actual := &components.Board{
		Cells: []*components.Cell{
			{Mark: NoMark},
			{Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}},
		Size: 9,
	}
	var i uint8
	for i = 0; i < size; i++ {
		if actual.Cells[i].Mark != expected.Board.Cells[i].Mark {
			t.Error("Error")
		}
	}
}

func TestPutMarkInPosition(t *testing.T) {
	expected := components.CreateBoard(3)
	simpleErr := expected.Cells[0].SetMark("X")
	if simpleErr != nil {
		t.Error(simpleErr)
	}
	actual := &BoardService{
		Board: components.CreateBoard(3),
	}
	myPlayer := &components.Player{
		Name: "golang",
		Mark: "X",
	}
	err := actual.PutMarkInPosition(myPlayer, 0)
	if err != nil {
		t.Error("Error in putting mark")
	}
	if actual.Board.Cells[0].Mark != expected.Cells[0].Mark {
		t.Error("Not working!")
	}
}

func TestBoardIsFull(t *testing.T) {
	myArray := make([]*components.Cell, 9)
	//Initially all cells will be blank:
	var i uint8
	for ; i < 9; i++ {
		myArray[i] = components.NewCell()
		myArray[i].SetMark("X")
	}
	//We have filled up every cell to be marked as "X" for testing puropose, now the board is full
	myBoard := &components.Board{
		Cells: myArray,
		Size:  3,
	}
	service := &BoardService{
		Board: myBoard,
	}
	//Checking if the function works!
	if !service.CheckBoardIsFull() {
		t.Error("Not working!")
	}
}

//Testing if the PrintBoard function works correctly and is able to print the board properly
func TestPrintBoard(t *testing.T) {
	myBoard := components.CreateBoard(3)
	boardService := NewBoardService(myBoard)
	actual := boardService.PrintBoard()
	expected := "- - - \n- - - \n- - - \n"
	if actual != expected {
		t.Error("not working")
		fmt.Println(expected)
		fmt.Println(actual)
	}
}
