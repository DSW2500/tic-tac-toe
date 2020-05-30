package services

import (
	"fmt"
	"math"
	"ticgame/components"
)

//ResultService :
type ResultService struct {
	BoardService *BoardService
	Result       string
}

//NewResultService :
func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{
		BoardService: boardService,
		Result:       "In Process",
	}
}

//GetResult : Public function which will be used and encapsulates all the other functions used for evaluating result
func (resultService *ResultService) GetResult(player *components.Player) string {
	resultService.Result = "The game is still in Process"
	res, winner := resultService.evaluateResult(player)
	if res {
		if winner == nil {
			resultService.Result = "It was a draw!"
		}
		resultService.Result = fmt.Sprintf("We have a winner\nPlayer : %s, Mark: %s has won!\n", winner.Name, winner.Mark)
	}

	return resultService.Result
}

func (resultService *ResultService) evaluateResult(player *components.Player) (bool, *components.Player) {
	if checkRows(resultService, player) {
		return true, player
	} else if checkColumns(resultService, player) {
		return true, player
	} else if checkDiagonal(resultService, player) {
		return true, player
	} else if resultService.BoardService.CheckBoardIsFull() {
		return true, nil
	}
	return false, nil
}

func checkRows(b *ResultService, player *components.Player) bool {
	//We check row-wise and iterate through cells
	flag := true
	size := int(b.BoardService.Board.Size)
	actualSize := int(math.Sqrt(float64(size)))
	for i := 0; i < (actualSize * actualSize); i = i + actualSize {
		flag = true
		for j := i; j < (i + actualSize); j++ {
			if j == size {
				break
			}
			if b.BoardService.Board.Cells[j].GetMark() != player.Mark {
				flag = false
			}

		}
		if flag {
			return flag
		}
	}
	return flag
}

func checkColumns(b *ResultService, player *components.Player) bool {
	flag := true
	size := int(b.BoardService.Board.Size)
	actualSize := int(math.Sqrt(float64(size)))
	for i := 0; i < actualSize; i++ {
		flag = true
		for j := i; j < (size); j = j + actualSize {
			if j == size {
				break
			}
			if b.BoardService.Board.Cells[j].GetMark() != player.Mark {
				flag = false
			}
		}
		if flag {
			return flag
		}
	}
	return flag
}
func checkDiagonal(b *ResultService, player *components.Player) bool {
	flag := true
	size := int(b.BoardService.Board.Size)
	actualSize := int(math.Sqrt(float64(size)))
	for i := 0; i < actualSize; i++ {
		if i == size {
			break
		}
		if b.BoardService.Board.Cells[actualSize*i+i].GetMark() != player.Mark {
			flag = false
		}
	}
	if flag {
		return flag
	}
	flag = true
	for i := 0; i < actualSize; i++ {
		if i == size {
			break
		}
		index := (actualSize * i) + (actualSize - 1 - i)
		if b.BoardService.Board.Cells[index].GetMark() != player.Mark {
			flag = false
		}
	}
	return flag
}
