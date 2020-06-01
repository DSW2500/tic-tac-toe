package main

import (
	"fmt"
	"ticgame/components"
	"ticgame/services"
)

func main() {
	//1. Rules are displayed
	services.DisplayRules()
	//2. Player names and marks are taken
	var player1 *components.Player
	var player2 *components.Player
	player1 = new(components.Player) // We will have two Players.
	player2 = new(components.Player)
	player1 = services.GetUserData(player1)
	player2 = services.GetUserData(player2)

	//3. Size of the board is taken from user
	size := services.GetBoardSize()
	//Total number of moves calculated using size
	//totalSize := int(size * size)
	//4. Board is created
	myBoard := components.CreateBoard(size)
	boardService := services.NewBoardService(myBoard)
	resultService := services.NewResultService(boardService)
	gameService := services.NewGameService(resultService)
	gameService.Players[0] = player1
	gameService.Players[1] = player2
	services.DisplayUserData(gameService.Players[0])
	services.DisplayUserData(gameService.Players[1])
	//Show initial board:
	fmt.Printf("The board looks like this\n%s", gameService.Result.BoardService.PrintBoard())
	//Match begins!
	gameService.SequencePlay(size)
}
