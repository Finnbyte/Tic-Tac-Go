package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"tictacgo/funcs"
	"tictacgo/locales"
	"tictacgo/utils"
	"time"
)

var (
	err          error
	Gamemode     string
	Side         string
	ComputerSide string
	cmd          string                                // Holds commands given in stdin
	Board        [][]string    = funcs.GenerateBoard() // Holds generated board; contents change during app's runtime
	lang         locales.Texts = locales.EngLang       // FiLang is finnish --- EngLang is english
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	/* SIDES */
	Side = string(funcs.GetSide()) // X or O
	ComputerSide = map[string]string{"X": "O", "O": "X"}[Side]

	/* INITIAL WELCOMING */
	fmt.Println(lang.Greeting)
	fmt.Println("\nWe have decided your side to be: " + Side + "\n")
	/* --------------------- */

	/* SETTING GAMEMODE */
	fmt.Println("Would you like to play:\n\t1 - Against a human locally?" +
		"\n\t2 - Against the computer?")
	for {
		fmt.Printf("\nUse \"help\" command to know what you can do!\n")
		utils.ScanlnWithPrompt(">>> ", &Gamemode)
		switch strings.ToLower(Gamemode) {
		case "1":
			utils.ClearOutput()
			versus(&Board)
		case "2":
			utils.ClearOutput()
			computer(&Board)
		case "help":
			fmt.Println("juu")
		case "exit":
			fmt.Println(lang.Exit)
			os.Exit(0)
		}
	}
}
