package services

import (
	"fmt"
	"ticgame/components"
)

//GameService :
type GameService struct {
	Result *ResultService
}

//NewGameService : New instance of GameService
func NewGameService(resultService *ResultService) *GameService {
	return &GameService{
		Result: resultService,
	}
}

//Play :
func (game *GameService) Play(moves int, players *components.Player) (int, bool) {
	//Accepting position from the user
	size := game.Result.BoardService.Board.Size
	//actualSize := int(math.Sqrt(float64(size)))

	position := game.getInput()
	err := game.Result.BoardService.PutMarkInPosition(players, position)
	if err != nil {
		fmt.Println(err)
		game.Play(moves, players)
	}
	fmt.Println("\nThe board currently looks like this:")
	fmt.Print(game.Result.BoardService.PrintBoard())
	if moves == int(size)-1 {
		res := game.Result.GetResult(players)
		if res != "The game is still in Process" {
			//fmt.Println(res)
			return moves, true
		}
	}
	moves++
	if moves >= 5 {
		res := game.Result.GetResult(players)
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
func DisplayUserData(players *components.Player) {

	fmt.Println("Player :", players.Name, " will play with the Mark: ", players.Mark)

}
