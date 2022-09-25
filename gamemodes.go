package main

import (
	"fmt"
	"strings"
	"tictacgo/funcs"
	"tictacgo/locales"
	"tictacgo/utils"
)

var (
	gameCmd string
	x       int8
	y       int8
)

func computer(board *[][]string) {
	for {
		fmt.Println(locales.EngLang.Prompt)
		utils.BufioWithPrompt(">>> ", &gameCmd)

		// Check if wants to exit
		if strings.ToLower(gameCmd) == "exit" {
			break
		}

		// Make input into x and y coords used for placing the symbol
		x, y, err = funcs.InterpretMoves(gameCmd)
		if err != nil {
			fmt.Printf("%s\n\n", err)
			continue
		}

		// Placing the symbol over coord
		if err = funcs.AssignMoves(board, Side, x, y); err != nil {
			fmt.Printf("%s\n\n", err)
			continue
		}

		// Computer makes a move on you
		funcs.ComputerMove(board, ComputerSide)

		// Show the board
		funcs.DisplayBoard(*board)

		// Check if player has won the game
		if won := funcs.HasWon(board, Side); won == true {
			fmt.Println(lang.Won)
			break
		}

		// Check if computer has won the game
		if won := funcs.HasWon(board, ComputerSide); won == true {
			fmt.Println(lang.Lost)
			break
		}

		// Check if game has come to a draw
		if draw := funcs.IsDraw(board); draw == true {
			fmt.Println(lang.Draw)
			break
		}
	}
}

func versus(board *[][]string) {
	for {
		fmt.Println(locales.EngLang.Prompt)
		utils.BufioWithPrompt(">>> ", &gameCmd)

		// Check if wants to exit
		if strings.ToLower(gameCmd) == "exit" {
			break
		}
	}
}
