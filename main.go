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
	player1 = new(components.Player) // We will have two players.
	player2 = new(components.Player)
	player1 = getUserData(player1)
	player2 = getUserData(player2)
	services.DisplayUserData(player1)
	services.DisplayUserData(player2)
	//3. Size of the board is taken from user
	size := getBoardSize()
	//Total number of moves calculated using size
	//totalSize := int(size * size)
	//4. Board is created
	myBoard := components.CreateBoard(size)
	boardService := services.NewBoardService(myBoard)
	resultService := services.NewResultService(boardService)
	gameService := services.NewGameService(resultService)
	//Show initial board:
	fmt.Printf("The board looks like this\n%s", gameService.Result.BoardService.PrintBoard())
	//Match begins!
	var result string
	flag := false
	moves := 0
	for i := 0; i <= int(size+1); i++ {

		if !flag {
			moves, flag = gameService.Play(moves, player1)
			if flag {
				result = gameService.Result.GetResult(player1)
				break
			}
			if i == int(size+1) {
				break
			}
			moves, flag = gameService.Play(moves, player2)
			if flag {
				result = gameService.Result.GetResult(player2)
				break
			}
		}

	}
	fmt.Println(result)

}

//getUserData :
func getUserData(player *components.Player) *components.Player {
	fmt.Println("User, please enter your name ")
	_, err := fmt.Scanf("%s", &player.Name)
	if err != nil {
		fmt.Println("Not a string! Enter again!")
		getUserData(player)
	}
	fmt.Println("User , please enter your mark! MUST BE 1 CHARACTER ONLY!!")
	_, err = fmt.Scanf("%s", &player.Mark)
	if err != nil {
		fmt.Println("Not a character, type again again!")
		getUserData(player)
	} else if len(player.Mark) > 1 {
		fmt.Println("Character length greater than 1! Enter again!!")
		getUserData(player)
	}
	return player
}

//getBoardSize :
func getBoardSize() uint8 {
	var size int
	fmt.Println("Please enter your board size! You can play with 3X3 or 4X4 for a quick game! Enter a single number")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println("Not an integer, try again!")
		getBoardSize()
	}
	if size > 4 || size < 3 {
		fmt.Println("Board size is not feasible. Try again!")
		getBoardSize()
	}
	return uint8(size)
}
