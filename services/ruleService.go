package services

import "fmt"

//DisplayRules :
func DisplayRules() {
	fmt.Println("\n\t\t\t\tWelcome to Tic Tac Toe!")
	fmt.Println("Game Rules:\n1. A user can win in 3 ways: joining three points diagonally, horizontally or vertically\n2. You will be shown the game's progress after every move\n3. To choose a square to mark on the board, simply type the index of that spot when prompted.")
	//Starting the game now:
	fmt.Print("\n\nThe Game begins now!\n\n")
}
