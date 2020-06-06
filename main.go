package main

import (
	"errors"
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
	player1 = getUserData(player1)
	player2 = getUserData(player2)
	//3. Size of the board is taken from user
	size := getBoardSize()
	//Total number of moves calculated using size
	//totalSize := int(size * size)
	//4. Board is created
	myBoard := components.CreateBoard(size)
	boardService := services.NewBoardService(myBoard)
	resultService := services.NewResultService(boardService)
	gameService := services.NewGameService(resultService)
	gameService.Players[0] = player1
	gameService.Players[1] = player2
	displayUserData(gameService.Players[0])
	displayUserData(gameService.Players[1])
	//Show initial board:
	fmt.Printf("The board looks like this\n%s", gameService.Result.BoardService.PrintBoard())
	//Match begins!

	var result string
	flag := false

	for i := 0; i <= int(size+1); i++ {
		var err error
		position := getInput()
		if !flag {
			flag, err = gameService.Play(position)
			if err == errors.New("Already marked cell") {
				position = getInput()
				gameService.Play(position)
			} else {
				fmt.Println(gameService.Result.BoardService.PrintBoard())
			}
			if flag {
				result = gameService.Result.GetResult(gameService.Players[0], position)
				break
			}
			if i == int(size+1) {
				break
			}
			position = getInput()
			flag, err = gameService.Play(position)
			if err != nil {
				position = getInput()
				gameService.Play(position)
			} else {
				fmt.Println(gameService.Result.BoardService.PrintBoard())
			}
			if flag {
				result = gameService.Result.GetResult(gameService.Players[1], position)
				break
			}

		}

	}
	fmt.Println(result)

}

//getInput :
func getInput() uint8 {
	var position int
	fmt.Println("Please enter the index")
	_, err := fmt.Scanf("%d", &position)
	if err != nil || position < 0 || position > 8 {
		fmt.Printf("Wrong index/not an integer between 0 and specified size\n")
		getInput()
	}
	return uint8(position)
}

//displayUserData :
func displayUserData(Players *components.Player) {
	fmt.Println("Player :", Players.Name, " will play with the Mark: ", Players.Mark)
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

//GetBoardSize :
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
