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
func (resultService *ResultService) GetResult(player *components.Player, index uint8) string {
	resultService.Result = "The game is still in Process"
	res, winner := resultService.evaluateResult(player, index)
	if res {
		resultService.Result = fmt.Sprintf("We have a winner\nPlayer : %s, Mark: %s has won!\n", winner.Name, winner.Mark)
	}
	if resultService.BoardService.CheckBoardIsFull() {
		resultService.Result = "It was a draw!"
	}

	return resultService.Result
}

func (resultService *ResultService) evaluateResult(player *components.Player, index uint8) (bool, *components.Player) {
	if resultService.checkRows(player, index) || resultService.checkColumns(player, index) || resultService.checkDiagonal(player) {
		return true, player
	}
	return false, nil
}

//checkRows :

func (rs *ResultService) checkRows(player *components.Player, index uint8) bool {
	cellArr := rs.BoardService.Board.Cells
	root := float64(rs.BoardService.Board.Size)
	limit := uint8(math.Sqrt(root))
	row := index / limit
	firstIndex := row * limit
	for i := firstIndex; i < firstIndex+limit; i++ {
		if cellArr[i].GetMark() != player.Mark {
			return false
		}
	}
	return true
}

// func checkRows(b *ResultService, player *components.Player) bool {
// 	//We check row-wise and iterate through cells
// 	flag := true
// 	size := int(b.BoardService.Board.Size)
// 	actualSize := int(math.Sqrt(float64(size)))
// 	for i := 0; i < (actualSize * actualSize); i = i + actualSize {
// 		flag = true
// 		for j := i; j < (i + actualSize); j++ {
// 			if j == size {
// 				break
// 			}
// 			if b.BoardService.Board.Cells[j].GetMark() != player.Mark {
// 				flag = false
// 			}

// 		}
// 		if flag {
// 			return flag
// 		}
// 	}
// 	return flag
// }
//checkColumns :
func (rs *ResultService) checkColumns(player *components.Player, index uint8) bool {
	cellArr := rs.BoardService.Board.Cells
	root := float64(rs.BoardService.Board.Size)
	limit := uint8(math.Sqrt(root))
	column := index % limit
	firstIndex := column * limit
	for i := firstIndex; i < uint8(root); i += limit {
		if cellArr[i].GetMark() != player.Mark {
			return false
		}
	}
	return true
}

// func checkColumns(b *ResultService, player *components.Player) bool {
// 	flag := true
// 	size := int(b.BoardService.Board.Size)
// 	actualSize := int(math.Sqrt(float64(size)))
// 	for i := 0; i < actualSize; i++ {
// 		flag = true
// 		for j := i; j < (size); j = j + actualSize {
// 			if j == size {
// 				break
// 			}
// 			if b.BoardService.Board.Cells[j].GetMark() != player.Mark {
// 				flag = false
// 			}
// 		}
// 		if flag {
// 			return flag
// 		}
// 	}
// 	return flag
// }
//checkDiagonals :
// func (rs *ResultService) checkDiagonals(player *components.Player, index uint8) bool {
// 	cellArr := rs.BoardService.Board.Cells
// 	root := float64(rs.BoardService.Board.Size)
// 	limit := uint8(math.Sqrt(root))
// 	column := index % limit
// 	firstIndex := column * limit
// 	for i := firstIndex; i < root; i += (limit + 1) {
// 		if cellArr[i].GetMark() != player.Mark {
// 			return false
// 		}
// 	}
// 	return true
// }

func (b *ResultService) checkDiagonal(player *components.Player) bool {
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
