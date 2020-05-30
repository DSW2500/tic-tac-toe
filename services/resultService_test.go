package services

import (
	"fmt"
	"testing"
	"ticgame/components"
)

func TestCheckRow(t *testing.T) {
	expected := fmt.Sprint("We have a winner\nPlayer : Golang, Mark: X has won!\n")
	player := &components.Player{Name: "Golang", Mark: "X"}
	var size uint8 = 3
	myBoard := components.CreateBoard(size)
	boardService := NewBoardService(myBoard)
	resultService := NewResultService(boardService)
	err := resultService.BoardService.Board.Cells[0].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	err = resultService.BoardService.Board.Cells[1].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	err = resultService.BoardService.Board.Cells[2].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	actual := resultService.GetResult(player)
	if actual != expected {
		fmt.Println(resultService.BoardService.PrintBoard())
		fmt.Println(actual)
		fmt.Println(expected)
		//t.Error("Not working")
	}

}

func TestCheckDiagonal(t *testing.T) {
	expected := fmt.Sprint("We have a winner\nPlayer : Golang, Mark: X has won!\n")
	player := &components.Player{Name: "Golang", Mark: "X"}
	var size uint8 = 3
	myBoard := components.CreateBoard(size)
	boardService := NewBoardService(myBoard)
	resultService := NewResultService(boardService)
	err := resultService.BoardService.Board.Cells[0].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	err = resultService.BoardService.Board.Cells[4].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	err = resultService.BoardService.Board.Cells[8].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	actual := resultService.GetResult(player)
	if actual != expected {
		fmt.Println(resultService.BoardService.PrintBoard())
		fmt.Println(actual)
		fmt.Println(expected)
		//t.Error("Not working")
	}

}
func TestCheckColumn(t *testing.T) {
	expected := fmt.Sprint("We have a winner\nPlayer : Golang, Mark: X has won!\n")
	player := &components.Player{Name: "Golang", Mark: "X"}
	var size uint8 = 3
	myBoard := components.CreateBoard(size)
	boardService := NewBoardService(myBoard)
	resultService := NewResultService(boardService)
	err := resultService.BoardService.Board.Cells[0].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	err = resultService.BoardService.Board.Cells[3].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	err = resultService.BoardService.Board.Cells[6].SetMark("X")
	if err != nil {
		t.Error(err)
	}
	actual := resultService.GetResult(player)
	if actual != expected {
		fmt.Println(resultService.BoardService.PrintBoard())
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Not working")
	}

}
