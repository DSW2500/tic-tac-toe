package services

import (
	"errors"
	"fmt"
	"math"
	"ticgame/components"
)

//GameService :
type GameService struct {
	Result  *ResultService
	Players [2]*components.Player
}

//Global variable : No of moves.
var moves int = 0

//NewGameService : New instance of GameService
func NewGameService(resultService *ResultService) *GameService {
	return &GameService{
		Result: resultService,
		Players: [2]*components.Player{
			{
				Name: "P1",
				Mark: "X",
			}, {
				Name: "P2",
				Mark: "O",
			},
		},
	}
}

//Play :
func (game *GameService) Play(position uint8) (bool, error) {
	var Players *components.Player
	//Accepting position from the user
	size := game.Result.BoardService.Board.Size
	if moves%2 == 0 {
		Players = game.Players[0]
	} else {
		Players = game.Players[1]
	}

	actualSize := int(math.Sqrt(float64(size)))
	err := game.Result.BoardService.PutMarkInPosition(Players, position)
	if err != nil {
		fmt.Println(err)
		return true, errors.New("Already marked cell")
	}
	if moves == int(size)-1 {
		res := game.Result.GetResult(Players)
		if res != "The game is still in Process" {
			//fmt.Println(res)
			return true, nil
		}
	}
	moves++
	if moves >= (2*actualSize - 1) {
		res := game.Result.GetResult(Players)
		if res != "The game is still in Process" {
			//fmt.Println(res)
			return true, nil
		}
	}

	return false, nil
}
