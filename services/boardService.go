package services

import (
	"errors"
	"math"
	"ticgame/components"
)

//BoardService :
type BoardService struct {
	Board *components.Board
}

//NewBoardService : Returns a board which will place a mark in the specified position
func NewBoardService(board *components.Board) *BoardService {
	//returns BoardInstance
	return &BoardService{
		Board: board,
	}
}

//PutMarkInPosition :
func (boardService *BoardService) PutMarkInPosition(player *components.Player, position uint8) error {
	if boardService.CheckBoardIsFull() {
		return errors.New("Board is full, so no mark can be placed")
	}
	if boardService.Board.Cells[position].GetMark() == components.OMark || boardService.Board.Cells[position].GetMark() == components.XMark {
		return errors.New("Cell has already been marked, please try another cell")
	}
	if boardService.Board.Cells[position].GetMark() == components.NoMark {
		boardService.Board.Cells[position].SetMark(player.Mark)
	}
	return nil
}

//PrintBoard : Returns a slice of Cells which can then be printed in a loop
func (boardService *BoardService) PrintBoard() string {
	size := int(boardService.Board.Size)
	actualSize := int(math.Sqrt(float64(size)))
	str := ""
	for i := 0; i < size; i++ {
		//fmt.Println(resultService.BoardService.Board.Cells[i].GetMark())
		str += boardService.Board.Cells[i].GetMark()
		str += " "
		if (i % actualSize) == ((actualSize) - 1) {
			str += "\n"
		}
	}
	return str
}

//CheckBoardIsFull : checks if the board is full
func (boardService *BoardService) CheckBoardIsFull() bool {
	//	flag := false
	var count uint8 = 0
	var i uint8
	//Loops through all cells and checks for NoMark, if at the end of loop, all elements have been marked, then board is full
	for i = 0; i < boardService.Board.Size; i++ {
		if boardService.Board.Cells[i].GetMark() != components.NoMark {
			// flag = true
			count++
		}
	}
	if count == boardService.Board.Size {
		return true
	}
	return false
}
