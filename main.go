package main

import (
	"bufio"
	"fmt"
	"os"
)

const XMove = "x"
const OMove = "o"
const rows = 3
const columns = 3

var player1 map[string]string
var player2 map[string]string

type BoardField struct {
	x, y int
	sign string
}

type Board [rows][columns]BoardField


func getPlayerName(player int) string {
	fmt.Printf("Player %v enter your name: ", player)
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	text := scanner.Text()

	return text
}

func getMove(playerName string) string {
	fmt.Print("")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

func printBoard(gameBoard Board) {

	for row := 0; row < rows; row++ { 
		for horizBorder := 0; horizBorder < (2 * columns + 1); horizBorder++ {
			fmt.Print("-")
		}

		fmt.Println()

		for column := 0; column < columns; column++ {
			fmt.Print("|")

			if gameBoard[row][column].sign == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(gameBoard[row][column].sign)
			}
			
			
			if column == columns -1 {
				fmt.Print("|")
			}
		}

		fmt.Println()

		if row == rows - 1 {
			for horizBorder := 0; horizBorder < (2 * columns + 1); horizBorder++ {
				fmt.Print("-")
			}
		}
	}
}

func isGameOver(gameBoard Board) {
	// Check rows
	
}

func main() {
	// player1Name := getPlayerName(1)
	// player2Name := getPlayerName(2)

	// player1 = make(map[string]string)
	// player1[player1Name] = XMove

	// player2 = make(map[string]string)
	// player2[player2Name] = OMove

	// gameBoard := Board{
	// 	{{x: 0, y: 0, sign: "o"}, {x: 1, y: 0, sign: "o"}, {x: 2, y: 0, sign: "o"}}, 
	// 	{{x: 0, y: 1, sign: "o"}, {x: 1, y: 1, sign: "o"}, {x: 2, y: 1, sign: "o"}}, 
	// 	{{x: 0, y: 2, sign: "o"}, {x: 1, y: 2, sign: "o"}, {x: 2, y: 2, sign: "o"}},
	// }

	var gameBoard Board

	printBoard(gameBoard)
}