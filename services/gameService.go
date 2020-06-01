package services

import (
	"fmt"
	"math"
	"ticgame/components"
)

//GameService :
type GameService struct {
	Result  *ResultService
	Players [2]*components.Player
}

//NewGameService : New instance of GameService
func NewGameService(resultService *ResultService) *GameService {
	return &GameService{
		Result: resultService,
	}
}
func (gameService *GameService) SequencePlay(size uint8) {
	var result string
	flag := false
	moves := 0
	for i := 0; i <= int(size+1); i++ {

		if !flag {
			moves, flag = gameService.Play(moves, gameService.Players[0])
			gameService.Result.BoardService.PrintBoard()
			if flag {
				result = gameService.Result.GetResult(gameService.Players[0])
				break
			}
			if i == int(size+1) {
				break
			}
			moves, flag = gameService.Play(moves, gameService.Players[1])
			gameService.Result.BoardService.PrintBoard()
			if flag {
				result = gameService.Result.GetResult(gameService.Players[1])
				break
			}

		}

	}
	fmt.Println(result)
}

//Play :
func (game *GameService) Play(moves int, Players *components.Player) (int, bool) {
	//Accepting position from the user
	size := game.Result.BoardService.Board.Size
	actualSize := int(math.Sqrt(float64(size)))

	position := game.getInput()
	err := game.Result.BoardService.PutMarkInPosition(Players, position)
	if err != nil {
		fmt.Println(err)
		game.Play(moves, Players)
	}
	if moves == int(size)-1 {
		res := game.Result.GetResult(Players)
		if res != "The game is still in Process" {
			//fmt.Println(res)
			return moves, true
		}
	}
	moves++
	if moves >= (2*actualSize - 1) {
		res := game.Result.GetResult(Players)
		if res != "The game is still in Process" {
			//fmt.Println(res)
			return moves, true
		}
	}

	return moves, false
}

//GetInput :
func (game *GameService) getInput() uint8 {
	size := int(game.Result.BoardService.Board.Size)
	var position int
	fmt.Println("Please enter your position: ")
	_, err := fmt.Scanf("%d", &position)
	if err != nil {
		fmt.Printf("Wrong index/not an integer between 0 and %d\n", size)
		game.getInput()
	}
	if position < 0 || position > 8 {
		fmt.Printf("Wrong index/not an integer between 0 and %d\n", size)
		game.getInput()
	}
	return uint8(position)
}

//DisplayUserData :
func DisplayUserData(Players *components.Player) {

	fmt.Println("Player :", Players.Name, " will play with the Mark: ", Players.Mark)

}
